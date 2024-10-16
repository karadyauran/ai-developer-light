package service

import (
	"ai-dev-light/internal/config"
	"ai-dev-light/internal/model"
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	ApiKey string
}

func NewOpenAIService(config *config.Config) *OpenAIService {
	return &OpenAIService{
		ApiKey: config.OpenApiKey,
	}
}

func (oas *OpenAIService) SendRequest(openAIRequest model.OpenAIRequest) (string, error) {
	client := openai.NewClient(oas.ApiKey)
	req := openai.ChatCompletionRequest{
		Model:       openAIRequest.Model,
		Messages:    chatRecordsToOpenAIMessages(openAIRequest.Messages),
		MaxTokens:   openAIRequest.MaxTokens,
		Temperature: openAIRequest.Temperature,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func chatRecordsToOpenAIMessages(records []model.Message) []openai.ChatCompletionMessage {
	var messages []openai.ChatCompletionMessage
	for _, record := range records {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    record.Role,
			Content: record.Content,
		})
	}
	return messages
}
