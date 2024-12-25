package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("󰑮 Creating a new Go project")
		// TODO: Add the init command logic

	},
}
