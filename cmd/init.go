package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vg006/vgo/app"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		app.Init()
	},
}
