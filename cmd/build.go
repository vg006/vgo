package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "up",
	Short: "Update the vgo tool to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("󰑮 Building")
		_, err := exec.Command("go", "build").Output()
		if err != nil {
			cmd.Println(" Error : Failed to update the vgo tool", err)
			return
		}
		cmd.Println(" Built")
		cmd.Println("󰑮 Installing")
		_, err = exec.Command("go", "install").Output()
		if err != nil {
			cmd.Println(" Error : Failed to update the vgo tool", err)
			return
		}
		cmd.Println(" Installed")

	},
}
