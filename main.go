package main

import (
	"aicoder/cmd"
	"aicoder/pkg/config"
	"aicoder/pkg/openai"
)

func main() {
	config.GetConfig()
	cmd.Execute()
	openai.DisposeClient()
}
