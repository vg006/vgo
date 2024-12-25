package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vgo",
	Short: "A Go project scaffolding tool",
	Long:  `vgo is a Go project scaffolding tool that helps you to create a new Go project with a predefined structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Vgo application")
	},
}

// --- Executes the command ---
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
