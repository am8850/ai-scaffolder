package cmd

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	prompt   string
	filePath string
)

const aicoder = `
 █████╗ ██╗ ██████╗ ██████╗ ██████╗ ███████╗██████╗ 
██╔══██╗██║██╔════╝██╔═══██╗██╔══██╗██╔════╝██╔══██╗
███████║██║██║     ██║   ██║██║  ██║█████╗  ██████╔╝
██╔══██║██║██║     ██║   ██║██║  ██║██╔══╝  ██╔══██╗
██║  ██║██║╚██████╗╚██████╔╝██████╔╝███████╗██║  ██║
╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝`

var rootCmd = &cobra.Command{
	Use:   "aicoder",
	Short: "aocoder CLI tool to generate or refactor code",
	Long:  `aicoder a command-line tool to generate or refactor code using AI from a user's prompt`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow.Printf(aicoder)
		fmt.Println("\nUse 'aicoder --help' for more information about using the tool.")
		color.Yellow.Println("\nTry:")
		color.Cyan.Println("  aicoder code -p \"Create a Python FastAPI application to manage customer.\"")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

}
