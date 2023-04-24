package services

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	client *openai.Client
}

// NewChatGPTClient returns a new openai client
func NewChatGPTClient(ctx context.Context, token string) *ChatGPTClient {
	return &ChatGPTClient{
		client: openai.NewClient(token),
	}
}

// AskAi asks ChatGPT to provide a review of a raw string, intended to be
// a git diff.
func (c *ChatGPTClient) AskAi(raw *string) (string, error) {
	question := `Please provide a review of the following git diff.
				Is the code idiomatic?
				Does the code make sense?
				Could anything be improved?`

	resp, err := c.client.CreateChatCompletion(
		context.TODO(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question + *raw,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
