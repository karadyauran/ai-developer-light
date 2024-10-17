package model

type Request struct {
	Request     string
	Model       string
	MaxTokens   int
	Temperature float32
}
