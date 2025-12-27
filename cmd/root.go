package cmd

import (
	"os"

	"github.com/spf13/cobra"
	asset "github.com/vg006/vgo/internal/assets"
)

var rootCmd = &cobra.Command{
	Use:   "vgo",
	Short: "A Go project scaffolding tool",
	Long:  `vgo is a Go project scaffolding tool that helps you to create a new Go project with a predefined structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(asset.VgoLogo)
		cmd.Println(asset.Text.Render("Welcome to vgo!"))
		cmd.Println(asset.Text.Render("vgo is a Go project scaffolding tool.\n"))
		cmd.Println(asset.Text.Render("Use 'vgo --help' for more information."))
	},
}

func Execute() {
	rootCmd.AddCommand(initCmd, updateCmd, buildCmd)
	rootCmd.PersistentFlags().BoolP("accessible", "a", false, "Set accessibility for screen reader compatibility")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
