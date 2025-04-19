package scaffolder

import (
	"aicoder/pkg/config"
	"aicoder/pkg/console"
	"aicoder/pkg/openai"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
)

func createFolderIfNotExists(filePath string) error {
	dir := filepath.Dir(filePath)

	if dir == "." {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

func generateCodeFiles(prompt string) (*config.CodeFiles, error) {
	appConfig := config.GetConfig()
	messages := []config.Message{
		{Role: "system", Content: appConfig.CodeSystemPrompt},
		{Role: "user", Content: prompt},
	}

	jdata, err := openai.ChatCompletion(messages, appConfig.Model, 0.1)
	if err != nil {
		return nil, err
	}

	var codefiles config.CodeFiles
	err = json.Unmarshal([]byte(jdata), &codefiles)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v (payload: %s)", err, jdata)
	}

	return &codefiles, nil
}

func displayCodeFiles(codefiles *config.CodeFiles) {
	fmt.Print("Generated code:\n\n")
	for _, codefile := range codefiles.Files {
		color.Yellow.Println("File: " + codefile.Filepath)
		color.Cyan.Println(codefile.Code + "\n")
	}
}

func writeCodeFiles(codefiles *config.CodeFiles) error {
	for _, codefile := range codefiles.Files {
		if err := createFolderIfNotExists(codefile.Filepath); err != nil {
			return fmt.Errorf("error creating directory for %s: %w", codefile.Filepath, err)
		}

		if err := os.WriteFile(codefile.Filepath, []byte(codefile.Code), 0644); err != nil {
			return fmt.Errorf("error writing file %s: %w", codefile.Filepath, err)
		}
	}
	return nil
}

func Scaffold(prompt string) {
	codefiles, err := generateCodeFiles(prompt)
	if err != nil {
		fmt.Println("Unable to generate code:")
		color.Red.Println(err)
		return
	}

	displayCodeFiles(codefiles)

	if console.AskForConfirmation("Do you want to write files?") {
		if err := writeCodeFiles(codefiles); err != nil {
			color.Red.Println(err)
		}
	}
}

func Refactor(filePath string, prompt string) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		color.Red.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	appConfig := config.GetConfig()
	messages := []config.Message{
		{Role: "system", Content: appConfig.CodeSystemPrompt},
		{Role: "user", Content: fmt.Sprintf("Please refactor the following code according to this instruction: %s\n\nCode:\n```\n%s\n```", prompt, string(content))},
	}

	jdata, err := openai.ChatCompletion(messages, appConfig.Model, 0.1)
	if err != nil {
		color.Red.Printf("Error generating refactored code: %v\n", err)
		return
	}

	// Extract code from response
	// This is a simple extraction. You might want to implement more robust parsing
	// depending on how your OpenAI model is configured to respond.
	refactoredCode := jdata

	// Display the refactored code
	color.Yellow.Printf("Original file: %s\n\n", filePath)
	color.Cyan.Println(refactoredCode + "\n")

	// Ask for confirmation before writing
	if console.AskForConfirmation("Do you want to replace the file with the refactored code?") {
		err = os.WriteFile(filePath, []byte(refactoredCode), 0644)
		if err != nil {
			color.Red.Printf("Error writing to file %s: %v\n", filePath, err)
		} else {
			color.Green.Println("File successfully refactored!")
		}
	}
}
