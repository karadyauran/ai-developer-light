package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"ai-dev-light/internal/config"
	"ai-dev-light/internal/model"

	"github.com/sashabaranov/go-openai"
)

const (
	RoleSystem    = "system"
	RoleAssistant = "assistant"
	RoleUser      = "user"

	DefaultModel       = "gpt-4o"
	DefaultMaxTokens   = 1500
	DefaultTemperature = 0.7
)

type AppBuilder struct {
	Client          *openai.Client
	ContextMessages []openai.ChatCompletionMessage
	Prompts         map[string]string
}

func NewAppBuilderService(cfg *config.Config) (*AppBuilder, error) {
	scriptDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	promptsDir := filepath.Join(scriptDir, "internal", "service", "prompts")

	contextFiles := map[string]string{
		"general_settings": filepath.Join(promptsDir, "general_settings.txt"),
	}

	promptFiles := map[string]string{
		"app_idea_generator":    filepath.Join(promptsDir, "app_idea_generator.txt"),
		"app_name_generator":    filepath.Join(promptsDir, "app_name_generator.txt"),
		"give_structure_to_gpt": filepath.Join(promptsDir, "give_structure_to_gpt.txt"),
		"default_file_content":  filepath.Join(promptsDir, "default_file_content.txt"),
		"create_readme_file":    filepath.Join(promptsDir, "create_readme_file.txt"),
		"get_gpt_files":         filepath.Join(promptsDir, "get_gpt_files.txt"),
		"generate_code":         filepath.Join(promptsDir, "generate_code.txt"),
		//"create_commit_message": filepath.Join(promptsDir, "create_commit_message.txt"),
	}

	contextMessages, err := loadContext(contextFiles)
	if err != nil {
		return nil, fmt.Errorf("failed to load context: %w", err)
	}

	prompts, err := loadPrompts(promptFiles)
	if err != nil {
		return nil, fmt.Errorf("failed to load prompts: %w", err)
	}

	return &AppBuilder{
		Client:          openai.NewClient(cfg.OpenApiKey),
		ContextMessages: contextMessages,
		Prompts:         prompts,
	}, nil
}

func loadContext(files map[string]string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage
	for _, filePath := range files {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read context file %s: %w", filePath, err)
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    RoleSystem,
			Content: string(content),
		})
	}
	return messages, nil
}

func loadPrompts(files map[string]string) (map[string]string, error) {
	prompts := make(map[string]string)
	for key, filePath := range files {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read prompt file %s: %w", filePath, err)
		}
		prompts[key] = string(content)
	}
	return prompts, nil
}

func (ap *AppBuilder) sendChatGPTRequest(request model.Request) (string, error) {
	userMessage := openai.ChatCompletionMessage{
		Role:    RoleUser,
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
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}

	answer := resp.Choices[0].Message.Content

	assistantMessage := openai.ChatCompletionMessage{
		Role:    RoleAssistant,
		Content: answer,
	}

	ap.ContextMessages = append(ap.ContextMessages, userMessage, assistantMessage)

	return answer, nil
}

func (ap *AppBuilder) generateAppIdea() error {
	content, exists := ap.Prompts["app_idea_generator"]
	if !exists {
		return fmt.Errorf("prompt 'app_idea_generator' not found")
	}

	return ap.sendRequestWithoutResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   100,
		Temperature: DefaultTemperature,
	})
}

func (ap *AppBuilder) generateAppName() (string, error) {
	content, exists := ap.Prompts["app_name_generator"]
	if !exists {
		return "", fmt.Errorf("prompt 'app_name_generator' not found")
	}

	return ap.sendRequestWithResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   25,
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

func (ap *AppBuilder) sendStructureOfProject() error {
	content, exists := ap.Prompts["give_structure_to_gpt"]
	if !exists {
		return fmt.Errorf("prompt 'give_structure_to_gpt' not found")
	}

	return ap.sendRequestWithoutResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   50,
		Temperature: DefaultTemperature,
	})
}

func (ap *AppBuilder) sendDefaultFilesForStyle() error {
	content, exists := ap.Prompts["default_file_content"]
	if !exists {
		return fmt.Errorf("prompt 'default_file_content' not found")
	}

	return ap.sendRequestWithoutResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   50,
		Temperature: DefaultTemperature,
	})
}

func (ap *AppBuilder) createReadMeFile(projectName string) error {
	content, exists := ap.Prompts["create_readme_file"]
	if !exists {
		return fmt.Errorf("prompt 'create_readme_file' not found")
	}

	readmeContent, err := ap.sendRequestWithResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   1500,
		Temperature: 1.0,
	})
	if err != nil {
		return err
	}

	filePath := filepath.Join("generated_projects", projectName, "README.md")
	return writeFileContent(filePath, readmeContent)
}

func (ap *AppBuilder) generaFilesNamesForStructure() ([]string, error) {
	content, exists := ap.Prompts["get_gpt_files"]
	if !exists {
		return nil, fmt.Errorf("prompt 'get_gpt_files' not found")
	}

	files, err := ap.sendRequestWithResponse(model.Request{
		Request:     content,
		Model:       DefaultModel,
		MaxTokens:   500,
		Temperature: DefaultTemperature,
	})
	if err != nil {
		return nil, err
	}

	return convertToSlice(files), nil
}

func (ap *AppBuilder) touchFiles(files []string, projectName string) error {
	for _, file := range files {
		file = removeExtraSpaces(file)
		fmt.Println(file, " generating")
		scriptPath := filepath.Join("internal", "utils", "create-file.sh")
		cmd := exec.Command("bash", scriptPath, projectName, file)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to create file: %v, output: %s", err, string(output))
		}
		err = ap.generateCode(file, projectName)
		if err != nil {
			return fmt.Errorf("failed to create file: %v, output: %s", err, string(output))
		}

		time.Sleep(7 * time.Second)
	}
	return nil
}

func removeExtraSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

func (ap *AppBuilder) generateCode(filePath string, projectName string) error {
	content, exists := ap.Prompts["generate_code"]
	if !exists {
		return fmt.Errorf("prompt 'generate_code' not found")
	}

	code, err := ap.sendRequestWithResponse(model.Request{
		Request:     fmt.Sprintf(content, filePath),
		Model:       DefaultModel,
		MaxTokens:   5000,
		Temperature: DefaultTemperature,
	})

	if err != nil {
		return err
	}

	fullFilePath := filepath.Join("generated_projects", projectName, filePath)
	err = writeFileContent(fullFilePath, code)
	if err != nil {
		return err
	}

	return nil
}

func convertToSlice(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func (ap *AppBuilder) sendRequestWithResponse(request model.Request) (string, error) {
	return ap.sendChatGPTRequest(request)
}

func (ap *AppBuilder) sendRequestWithoutResponse(request model.Request) error {
	_, err := ap.sendChatGPTRequest(request)
	return err
}

func writeFileContent(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to file %s: %w", filePath, err)
	}

	return nil
}

func (ap *AppBuilder) BuildWithNoContext() error {
	if err := ap.generateAppIdea(); err != nil {
		return fmt.Errorf("failed to generate app idea: %w", err)
	}
	fmt.Println("Generated Idea")

	appName, err := ap.generateAppName()
	if err != nil {
		return fmt.Errorf("failed to generate app name: %w", err)
	}
	fmt.Println("App name is ", appName)

	if err := createBaseStructure(appName); err != nil {
		return fmt.Errorf("failed to create base structure: %w", err)
	}
	fmt.Println("Created base structure")

	if err := ap.sendStructureOfProject(); err != nil {
		return fmt.Errorf("failed to send structure of project: %w", err)
	}
	fmt.Println("Sent structure of project")

	if err := ap.sendDefaultFilesForStyle(); err != nil {
		return fmt.Errorf("failed to send default files for style: %w", err)
	}
	fmt.Println("Default files for style")

	if err := ap.createReadMeFile(appName); err != nil {
		return fmt.Errorf("failed to create README file: %w", err)
	}
	fmt.Println("Created ReadME file")

	fmt.Println("Starting generating files")
	filesList, err := ap.generaFilesNamesForStructure()
	if err != nil {
		return fmt.Errorf("failed to generate project structure: %w", err)
	}
	fmt.Println("Complete!")

	fmt.Println("Starting creating files")
	if err := ap.touchFiles(filesList, appName); err != nil {
		return fmt.Errorf("failed to touch files for structure: %w", err)
	}
	fmt.Println("Complete!")

	return nil
}
