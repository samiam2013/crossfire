package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/samiam2013/crossfire/history"
	log "github.com/sirupsen/logrus"
)

type ChatCompletionRequest struct {
	Model string `json:"model"`
	// Format   ResponseFormat `json:"response_format"`
	Messages []ChatMessage `json:"messages"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func (ccp ChatCompletionResponse) FirstContent() (string, error) {
	if len(ccp.Choices) == 0 {
		return "", fmt.Errorf("no choices in completion response")
	}
	return ccp.Choices[0].Message.Content, nil
}

type API struct {
	openAIKey string
}

func NewAPI(openAIKey string) *API {
	return &API{
		openAIKey: openAIKey,
	}
}

func (a *API) GetCompletion(userInput string, hist history.MessageHistory) (c ChatCompletionResponse, err error) {
	msgs := make([]ChatMessage, 0)
	msgs = append(msgs, ChatMessage{
		Role:    "system",
		Content: "you are a large language model used for debating ideas as concisely (short responses) as possible",
	})
	for _, msg := range hist {
		if msg.Author == history.AuthorOpenAI {
			msgs = append(msgs, ChatMessage{Role: "system", Content: msg.Content})
		} else {
			msgs = append(msgs, ChatMessage{Role: "user", Content: msg.Content})
		}
	}
	requestData := ChatCompletionRequest{
		Model: "gpt-4-turbo-preview",
		Messages: append(msgs, ChatMessage{
			Role:    "user",
			Content: userInput,
		}),
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return c, fmt.Errorf("error marshaling JSON: %w", err)
	}

	u, err := url.Parse("https://api.openai.com/v1/chat/completions")
	if err != nil {
		return c, fmt.Errorf("error parsing URL: %w", err)
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return c, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.openAIKey))

	// Initialize an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c, fmt.Errorf("error sending request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return c, fmt.Errorf("response status not OK: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		if err == io.EOF {
			return c, nil
		}
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	return c, nil
}

type Address struct {
	Number        string `json:"number"`
	StreetAddress string `json:"street_address"`
	Unit          string `json:"unit"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
}

func (a *API) GetUSAddressSorted(input string) (sa Address, err error) {
	goTypePrompt := "```go\n" +
		"type Address struct {\n" +
		"	Number 		  string    `json:\"number\"`\n" +
		"	StreetAddress string `json:\"street_address\"// add periods after abbreviations`\n" +
		"   Unit 		  string `json:\"unit\"`\n" +
		"	City          string `json:\"city\"`\n" +
		"	State         string `json:\"state\"`// 2-letter state code\n" +
		"	ZipCode       string `json:\"zip_code\"` // no local route codes\n" +
		"	Country       string `json:\"country\"` // 2-letter country code\n" +
		"}\n" +
		"```"

	// create a prompt with a go type for a US address and a field if it's foreign
	textPrompt := "please make a JSON response that will unmarshall this address `" + input + "` in this go struct format " +
		goTypePrompt + " and be liberal about normalizing to colloquial US address formatting"

	completion, err := a.GetCompletion(textPrompt, history.NewMessageHistory())
	if err != nil {
		return sa, fmt.Errorf("error getting completion: %v", err)
	}
	content := completion.Choices[0].Message.Content
	log.Printf("Completion: %s", content)

	if err := json.Unmarshal([]byte(content), &sa); err != nil {
		return sa, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return
}
