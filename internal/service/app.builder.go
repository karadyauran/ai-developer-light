package service

import (
	"ai-dev-light/internal/model"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	placeholderLanguage = "{language}"
	placeholderNFiles   = "{n_files}"
	placeholderExt      = "{ext}"
	placeholderFileName = "{file_name}"
)

type AppBuilder struct {
	OpenAIService *OpenAIService
	Prompts       map[string]string
	Languages     []string
	History       []model.Message
}

func NewAppBuilderService(openAIService *OpenAIService) (*AppBuilder, error) {
	prompts, err := loadPrompts()
	if err != nil {
		return nil, err
	}

	return &AppBuilder{
		OpenAIService: openAIService,
		Prompts:       prompts,
		Languages:     []string{"Go", "Python", "JavaScript"},
		History: []model.Message{
			{Role: "system", Content: "You are a senior software developer proficient in multiple programming languages."},
		},
	}, nil
}

func loadPrompts() (map[string]string, error) {
	prompts := make(map[string]string)
	scriptDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	promptFiles := map[string]string{
		"app_idea":        filepath.Join(scriptDir, "internal", "scripts", "app_idea_prompt.txt"),
		"file_structure":  filepath.Join(scriptDir, "internal", "scripts", "file_structure_prompt.txt"),
		"code_generation": filepath.Join(scriptDir, "internal", "scripts", "code_generation_prompt.txt"),
		"commit_message":  filepath.Join(scriptDir, "internal", "scripts", "commit_message_prompt.txt"),
	}

	for key, filePath := range promptFiles {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		prompts[key] = string(content)
	}
	return prompts, nil
}

func (ab *AppBuilder) GenerateUniqueAppName(baseName string) string {
	uniqueID := uuid.New().String()[:8]
	return fmt.Sprintf("%s_%s", baseName, uniqueID)
}

func (ab *AppBuilder) GenerateAppIdea(language string, nFiles int) (string, error) {
	promptTemplate, ok := ab.Prompts["app_idea"]
	if !ok {
		return "", errors.New("app_idea prompt not found")
	}
	prompt := strings.ReplaceAll(promptTemplate, placeholderLanguage, language)
	prompt = strings.ReplaceAll(prompt, placeholderNFiles, fmt.Sprintf("%d", nFiles))

	ab.History = append(ab.History, model.Message{Role: "user", Content: prompt})

	openAIRequest := model.OpenAIRequest{
		Model:       ab.OpenAIService.Model,
		MaxTokens:   150,
		Temperature: 0.7,
		Messages:    ab.History,
	}

	response, err := ab.OpenAIService.SendRequest(openAIRequest)
	if err != nil {
		return "", err
	}

	ab.History = append(ab.History, model.Message{Role: "assistant", Content: response})

	return response, nil
}

func (ab *AppBuilder) GenerateFileStructure(language string, nFiles int) ([]string, error) {
	promptTemplate, ok := ab.Prompts["file_structure"]
	if !ok {
		return nil, errors.New("file_structure prompt not found")
	}
	ext := ab.GetFileExtension(language)
	prompt := strings.ReplaceAll(promptTemplate, placeholderLanguage, language)
	prompt = strings.ReplaceAll(prompt, placeholderNFiles, fmt.Sprintf("%d", nFiles))
	prompt = strings.ReplaceAll(prompt, placeholderExt, ext)

	ab.History = append(ab.History, model.Message{Role: "user", Content: prompt})

	openAIRequest := model.OpenAIRequest{
		Model:       ab.OpenAIService.Model,
		MaxTokens:   50,
		Temperature: 0.7,
		Messages:    ab.History,
	}

	response, err := ab.OpenAIService.SendRequest(openAIRequest)
	if err != nil {
		return nil, err
	}

	ab.History = append(ab.History, model.Message{Role: "assistant", Content: response})

	// Process response into a slice of file names
	files := strings.Split(response, ",")
	for i, file := range files {
		files[i] = strings.TrimSpace(file)
	}

	return files, nil
}

func (ab *AppBuilder) GenerateCode(fileName, language string) (string, error) {
	promptTemplate, ok := ab.Prompts["code_generation"]
	if !ok {
		return "", errors.New("code_generation prompt not found")
	}
	prompt := strings.ReplaceAll(promptTemplate, placeholderFileName, fileName)
	prompt = strings.ReplaceAll(prompt, placeholderLanguage, language)

	ab.History = append(ab.History, model.Message{Role: "user", Content: prompt})

	openAIRequest := model.OpenAIRequest{
		Model:       ab.OpenAIService.Model,
		MaxTokens:   1000,
		Temperature: 0.7,
		Messages:    ab.History,
	}

	response, err := ab.OpenAIService.SendRequest(openAIRequest)
	if err != nil {
		return "", err
	}

	ab.History = append(ab.History, model.Message{Role: "assistant", Content: response})

	code := processCodeResponse(response)

	return code, nil
}

func processCodeResponse(response string) string {
	lines := strings.Split(response, "\n")
	if len(lines) > 2 {
		return strings.Join(lines[1:len(lines)-1], "\n")
	}
	return response
}

func (ab *AppBuilder) CreateCommitMessage(fileName string) (string, error) {
	promptTemplate, ok := ab.Prompts["commit_message"]
	if !ok {
		return "", errors.New("commit_message prompt not found")
	}
	prompt := strings.ReplaceAll(promptTemplate, placeholderFileName, fileName)

	ab.History = append(ab.History, model.Message{Role: "user", Content: prompt})

	openAIRequest := model.OpenAIRequest{
		Model:       ab.OpenAIService.Model,
		MaxTokens:   50,
		Temperature: 0.7,
		Messages:    ab.History,
	}

	response, err := ab.OpenAIService.SendRequest(openAIRequest)
	if err != nil {
		return "", err
	}

	ab.History = append(ab.History, model.Message{Role: "assistant", Content: response})

	return response, nil
}

func (ab *AppBuilder) GetFileExtension(language string) string {
	extensions := map[string]string{
		"Go":         "go",
		"Python":     "py",
		"JavaScript": "js",
	}
	if ext, exists := extensions[language]; exists {
		return ext
	}
	return "txt"
}

func (ab *AppBuilder) CreateRepository(appName string) error {
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}
	appDir := filepath.Join(baseDir, "projects", appName)
	err = os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		return err
	}
	return os.Chdir(appDir)
}
