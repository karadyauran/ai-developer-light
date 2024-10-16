package service

import (
	"ai-dev-light/internal/config"
	"ai-dev-light/internal/model"
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	ApiKey   string
	Model    string
	Messages []model.Message
}

func NewOpenAIService(cfg *config.Config, chatModel string) *OpenAIService {
	return &OpenAIService{
		ApiKey:   cfg.OpenApiKey,
		Model:    chatModel,
		Messages: []model.Message{},
	}
}

func (oas *OpenAIService) SendRequest(openAIRequest model.OpenAIRequest) (string, error) {
	oas.Messages = openAIRequest.Messages

	client := openai.NewClient(oas.ApiKey)
	req := openai.ChatCompletionRequest{
		Model:       openAIRequest.Model,
		Messages:    chatRecordsToOpenAIMessages(oas.Messages),
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

func (oas *OpenAIService) ResetMessages() {
	oas.Messages = []model.Message{}
}
