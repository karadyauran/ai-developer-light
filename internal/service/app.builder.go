package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"ai-dev-light/internal/config"
	"ai-dev-light/internal/model"
)

const (
	systemRole    = "system"
	assistantRole = "assistant"
	userRole      = "user"
)

type AppBuilder struct {
	Client          *openai.Client
	ContextMessages []openai.ChatCompletionMessage
	PromptFiles     map[string]string
}

func NewAppBuilderService(cfg *config.Config) *AppBuilder {
	scriptDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	promptsDir := filepath.Join(scriptDir, "internal", "service", "prompts")

	contextFiles := map[string]string{
		"general_settings": filepath.Join(promptsDir, "general_settings.txt"),
	}

	prompts := map[string]string{
		"app_idea_generator":    filepath.Join(promptsDir, "app_idea_generator.txt"),
		"app_name_generator":    filepath.Join(promptsDir, "app_name_generator.txt"),
		"give_structure_to_gpt": filepath.Join(promptsDir, "give_structure_to_gpt.txt"),
		"get_gpt_files":         filepath.Join(promptsDir, "get_gpt_files.txt"),
		"generate_code":         filepath.Join(promptsDir, "generate_code.txt"),
		"create_commit_message": filepath.Join(promptsDir, "create_commit_message.txt"),
	}

	contextMessages, err := loadContext(contextFiles)
	if err != nil {
		log.Fatal(err)
	}

	return &AppBuilder{
		Client:          openai.NewClient(cfg.OpenApiKey),
		ContextMessages: contextMessages,
		PromptFiles:     prompts,
	}
}

func loadContext(promptFiles map[string]string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	for _, filePath := range promptFiles {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    systemRole,
			Content: string(content),
		})
	}

	return messages, nil
}

func (ap *AppBuilder) sendChatGPTRequest(request model.Request) (string, error) {
	userMessage := openai.ChatCompletionMessage{
		Role:    userRole,
		Content: request.Request,
	}

	messages := append(ap.ContextMessages, userMessage)

	req := openai.ChatCompletionRequest{
		Model:       request.Model,
		Messages:    messages,
		MaxTokens:   request.MaxTokens,
		Temperature: request.Temperature,
	}

	resp, err := ap.Client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content

	assistantMessage := openai.ChatCompletionMessage{
		Role:    assistantRole,
		Content: answer,
	}

	// Update the context messages
	ap.ContextMessages = append(messages, assistantMessage)

	return answer, nil
}

func (ap *AppBuilder) sendRequestToChatGPTWithResponse(request model.Request) (string, error) {
	return ap.sendChatGPTRequest(request)
}

func (ap *AppBuilder) sendRequestToChatGPTWithoutResponse(request model.Request) error {
	_, err := ap.sendChatGPTRequest(request)
	return err
}

func getContentFromFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (ap *AppBuilder) generateAppIdea() error {
	content, err := getContentFromFile(ap.PromptFiles["app_idea_generator"])
	if err != nil {
		return err
	}

	return ap.sendRequestToChatGPTWithoutResponse(model.Request{
		Request:     content,
		Model:       "gpt-4o",
		MaxTokens:   10000,
		Temperature: 0.7,
	})
}

func (ap *AppBuilder) generateAppName() (string, error) {
	content, err := getContentFromFile(ap.PromptFiles["app_name_generator"])
	if err != nil {
		return "", err
	}

	return ap.sendRequestToChatGPTWithResponse(model.Request{
		Request:     content,
		Model:       "gpt-4o",
		MaxTokens:   50,
		Temperature: 1.0,
	})
}

func createBaseStructure(appName string) error {

	scriptPath := filepath.Join("internal", "utils", "generate-arch.sh")
	cmd := exec.Command("bash", scriptPath, appName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create base structure: %v, output: %s", err, string(output))
	}
	return nil
}

func (ap *AppBuilder) BuildWithNoContext() error {
	if err := ap.generateAppIdea(); err != nil {
		return err
	}

	appName, err := ap.generateAppName()
	if err != nil {
		return err
	}

	if err := createBaseStructure(appName); err != nil {
		return err
	}

	return nil
}
