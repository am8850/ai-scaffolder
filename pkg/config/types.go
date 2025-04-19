package config

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponsFormatType struct {
	Type string `json:"type"`
}

// Payload represents the JSON payload to be sent
type ChatRequest struct {
	Messages       []Message              `json:"messages"`
	Model          string                 `json:"model"`
	Temperature    float64                `json:"temperature"`
	ResponseFormat *ChatResponsFormatType `json:"response_format"`
}

// Response represents the JSON response from the API
type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// Command represents the command to be executed
type Command struct {
	Command     string   `json:"command"`
	Args        []string `json:"args"`
	Explanation string   `json:"explanation"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

type CodeFile struct {
	Filepath string `json:"filepath"`
	Code     string `json:"code"`
}

type CodeFiles struct {
	Files []CodeFile `json:"files"`
}

type SanitizerResponse struct {
	ReadabilityScore  int    `json:"readability_score"`
	ReadabilityReason string `json:"readability_reason"`
	CyclomaticScore   int    `json:"cyclomatic_score"`
	CyclomaticReason  string `json:"cyclomatic_reason"`
	ImprovedCode      string `json:"improved_code"`
}

type SystemPrompt struct {
	Command      string `json:"command"`
	SystemPrompt string `json:"system"`
}
