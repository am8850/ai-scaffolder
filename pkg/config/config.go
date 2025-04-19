package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Endpoint             string `json:"endpoint"`
	Key                  string `json:"key"`
	Model                string `json:"model"`
	Type                 string `json:"type"` // "openai" or "azure"
	CodeSystemPrompt     string `json:"code_system_prompt"`
	RefactorSystemPrompt string `json:"refactor_system_prompt"`
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton instance of the configuration
func GetConfig() *Config {
	var err error
	once.Do(func() {
		instance = &Config{}
		err = loadConfig()
		if err != nil {
			panic("Failed to load configuration: " + err.Error())
		}
		if instance.Type == "" {
			instance.Type = "openai"
		}
		if instance.Endpoint == "" || instance.Key == "" || instance.Model == "" || instance.CodeSystemPrompt == "" || instance.RefactorSystemPrompt == "" {
			panic("Missing required configuration fields: endpoint, key, or model in scai.json")
		}
	})
	return instance
}

// loadConfig reads the configuration from the scai.json file
func loadConfig() error {
	// Try to find configuration in the current working directory
	file, err := os.Open("./aicoder.json")
	if err != nil {
		// If not found, try from executable's directory
		execPath, execErr := os.Executable()
		if execErr != nil {
			return err // Return original error if we can't get executable path
		}

		execDir := filepath.Dir(execPath)
		file, err = os.Open(filepath.Join(execDir, "aicoder.json"))
		if err != nil {
			return err // Configuration file not found in either location
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&instance)
}
