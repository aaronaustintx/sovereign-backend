package ai

import (
    "context"

    openai "github.com/sashabaranov/go-openai"
)

type Client struct {
    api *openai.Client
}

func New(apiKey string) *Client {
    return &Client{
        api: openai.NewClient(apiKey),
    }
}

func (c *Client) Generate(model, prompt string) (string, error) {
    resp, err := c.api.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: model,
            Messages: []openai.ChatCompletionMessage{
                {Role: openai.ChatMessageRoleUser, Content: prompt},
            },
        },
    )
    if err != nil {
        return "", err
    }

    return resp.Choices[0].Message.Content, nil
}
