package anthropic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type API struct {
	anthropicKey string
}

func NewAPI(anthropicKey string) *API {
	return &API{
		anthropicKey: anthropicKey,
	}
}

/*
	curl https://api.anthropic.com/v1/messages \
	     --header "x-api-key: $ANTHROPIC_API_KEY" \
	     --header "anthropic-version: 2023-06-01" \
	     --header "content-type: application/json" \
	     --data \

	'{
	    "model": "claude-3-opus-20240229",
	    "max_tokens": 1024,
	    "messages": [
	        {"role": "user", "content": "Hello, world"}
	    ]
	}'
*/
type MessageRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

/*
	{
		"content": [
		  {
			"text": "Hi! My name is Claude.",
			"type": "text"
		  }
		],
		"id": "msg_013Zva2CMHLNnXjNJJKqJ2EF",
		"model": "claude-3-opus-20240229",
		"role": "assistant",
		"stop_reason": "end_turn",
		"stop_sequence": null,
		"type": "message",
		"usage": {
		  "input_tokens": 10,
		  "output_tokens": 25
		}
	  }
*/
type MessageResponse struct {
	Content []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
	ID           string `json:"id"`
	Model        string `json:"model"`
	Role         string `json:"role"`
	StopReason   string `json:"stop_reason"`
	StopSequence string `json:"stop_sequence"`
	Type         string `json:"type"`
	Usage        struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

func (a API) GetMessageResponse(prompt string) (mr MessageResponse, err error) {
	msg := MessageRequest{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 1024,
		Messages: []Message{{
			Role:    "user",
			Content: prompt,
		}},
	}
	marshalled, err := json.Marshal(msg)
	if err != nil {
		return mr, errors.Join(err, errors.New("failed to marshal message request"))
	}

	// build the message into an HTTP request
	req, err := http.NewRequest(http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewBuffer(marshalled))
	if err != nil {
		return mr, errors.Join(err, errors.New("failed to create request"))
	}
	req.Header.Set("x-api-key", a.anthropicKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return mr, errors.Join(err, errors.New("failed to send request"))
	}
	defer func() { _ = resp.Body.Close() }()

	byts, err := io.ReadAll(resp.Body)
	log.Infof("response body: %s", string(byts))
	if err != nil {
		return mr, errors.Join(err, errors.New("failed to read response body"))
	}
	if err := json.Unmarshal(byts, &mr); err != nil {
		return mr, errors.Join(err, errors.New("failed to unmarshal response"))
	}
	return mr, nil
}
