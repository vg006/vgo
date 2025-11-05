package cmd

import (
	"fmt"
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
		var (
			err        error
			flag       = true
			buildOut   []byte
			installOut []byte
		)

		_ = spinner.
			New().
			Title("Building").
			Action(func() {
				buildOut, err = exec.Command("go", "build").CombinedOutput()
				if err != nil {
					flag = false
				}

			}).
			Style(asset.Text).
			Accessible(false).
			Run()

		if flag {
			cmd.Println(asset.Text.Foreground(asset.Green).
				Render(fmt.Sprintf("%s Built vgo", asset.EmojiTick)))
		} else {
			cmd.Println(asset.Text.Foreground(asset.Red).
				Render(fmt.Sprintf("%s Error : Failed to build the vgo tool\n%v\n%s", asset.EmojiError, err, string(buildOut))))
			return
		}

		_ = spinner.
			New().
			Title("Installing").
			Action(func() {
				installOut, err = exec.Command("go", "install").CombinedOutput()
				if err != nil {
					flag = false
				}
			}).
			Style(asset.Text).
			Accessible(false).
			Run()

		if flag {
			cmd.Println(asset.Text.Foreground(asset.Green).
				Render(fmt.Sprintf("%s Installed vgo", asset.EmojiTick)))
		} else {
			cmd.Println(asset.Text.Foreground(asset.Red).
				Render(fmt.Sprintf("%s Error : Failed to update the vgo tool\n%v\n%s", asset.EmojiError, err, string(installOut))))
			return
		}
	},
}
