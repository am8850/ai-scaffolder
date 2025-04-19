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

		fmt.Println("Writing file:", codefile.Filepath)
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
