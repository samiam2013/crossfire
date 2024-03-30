package anthropic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/samiam2013/crossfire/pkg/history"
)

type API struct {
	anthropicKey string
}

func NewAPI(anthropicKey string) *API {
	return &API{
		anthropicKey: anthropicKey,
	}
}

type MessageRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

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

func (mr MessageResponse) FirstContent() (string, error) {
	if len(mr.Content) == 0 {
		return "", fmt.Errorf("no content (length 0) in message response")
	}
	return mr.Content[0].Text, nil
}

func (a API) GetMessageResponse(prompt string, hist history.MessageHistory) (mr MessageResponse, err error) {
	msgs := make([]Message, 0)
	msgs = append(msgs, Message{
		Role:    "user",
		Content: prompt,
	})
	for _, h := range hist {
		if h.Author == history.AuthorClaude {
			msgs = append(msgs, Message{
				Role:    "assistant",
				Content: h.Content,
			})
		} else {
			msgs = append(msgs, Message{
				Role:    "user",
				Content: h.Content,
			})
		}
	}

	msg := MessageRequest{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 1000,
		Messages: append(msgs, Message{
			Role: "assistant",
			Content: "Hi! My name is Claude. I'm a debate assistant. I will very very concisely (short messages) respond to a prompt." +
				"I might not agree with the side I'm arguing for, but I'll do my best to make a compelling argument:",
		}),
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
	if resp.StatusCode != http.StatusOK {
		return mr, fmt.Errorf("failure status code (!200): %d", resp.StatusCode)
	}
	if err != nil {
		return mr, errors.Join(err, errors.New("failed to read response body"))
	}
	if err := json.Unmarshal(byts, &mr); err != nil {
		return mr, errors.Join(err, errors.New("failed to unmarshal response"))
	}

	return mr, nil
}
