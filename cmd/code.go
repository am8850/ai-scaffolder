package cmd

import (
	"aicoder/pkg/scaffolder"
	"fmt"

	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:     "code",
	Aliases: []string{"co"},
	Short:   "Generate code from a prompt",
	Long:    `Scaffold new code from a prompt using AI`,
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Error: --prompt or -p flag is required")
		} else {
			scaffolder.Scaffold(prompt)
		}
	},
}

func init() {
	codeCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Prompt for the CLI")
	codeCmd.MarkPersistentFlagRequired("prompt")
	rootCmd.AddCommand(codeCmd)
}
