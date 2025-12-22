package cmd

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	asset "github.com/vg006/vgo/internal/assets"
)

var updateCmd = &cobra.Command{
	Use:   "up",
	Short: "Update the vgo tool to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		accessible, _ := cmd.Flags().GetBool("accessible")
		_ = spinner.
			New().
			Title("Updating vgo ...").
			Action(func() {
				_, err := exec.Command("go", "install", "github.com/vg006/vgo@latest").Output()
				if err != nil {
					cmd.Println(
						asset.Text.Foreground(asset.Red).
							Render(fmt.Sprintf("%s Error : Failed to update the vgo tool", asset.EmojiError)))
				}
				cmd.Println(
					asset.Text.Foreground(asset.Green).
						Render(fmt.Sprintf("%s Yupee! You are now up-to-date!", asset.EmojiSparkles)))
			}).
			Style(asset.Text).
			Accessible(accessible).
			Run()
	},
}
