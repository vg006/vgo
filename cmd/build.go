package cmd

import (
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	asset "github.com/vg006/vgo/internal/assets"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the vgo tool and install it",
	Run: func(cmd *cobra.Command, args []string) {
		_ = spinner.
			New().
			Title("Building vgo ...").
			Action(func() {
				_, err := exec.Command("go", "build").Output()
				if err != nil {
					cmd.Printf("%s Error : Failed to update the vgo tool", asset.EmojiError)
					return
				}
				cmd.Printf("%s Built", asset.EmojiTick)
			}).
			Style(asset.Text).
			Accessible(true).
			Run()

		_ = spinner.
			New().
			Title("Installing vgo ...").
			Action(func() {
				_, err := exec.Command("go", "install").Output()
				if err != nil {
					cmd.Printf("%s Error : Failed to update the vgo tool", asset.EmojiError)
					return
				}
				cmd.Printf("%s Installed", asset.EmojiTick)
			}).
			Style(asset.Text).
			Accessible(true).
			Run()
	},
}
