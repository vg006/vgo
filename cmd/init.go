package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	app "github.com/vg006/vgo/internal"
	asset "github.com/vg006/vgo/internal/assets"
	"github.com/vg006/vgo/internal/utils"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		var p app.Project

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&p.Name).
					Title("Project Name").
					Description("Enter the project name").
					Validate(utils.CheckValidProjectName),
			),
			huh.NewGroup(
				huh.NewInput().
					Value(&p.ModName).
					Title("Module Name").
					Description("Enter a name of the module\nPress <Enter> to use the same as project name").
					Validate(func(s string) error {
						if s == "" {
							p.ModName = p.Name
						}
						return nil
					}),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&p.FrameWork).
					Title("Framework").
					Description("Select the web framework for the project").
					Options(
						huh.NewOption("Go Standard Library", "stdlib"),
						huh.NewOption("Echo", "echo"),
						huh.NewOption("Fiber", "fiber"),
						huh.NewOption("Chi", "chi"),
						huh.NewOption("Gin", "gin"),
					),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&p.Database).
					Title("Database").
					Description("Select the database for the project").
					Options(
						huh.NewOption("None", "none"),
						huh.NewOption("PostgreSQL", "postgresql"),
						huh.NewOption("MySQL", "mysql"),
						huh.NewOption("SQLite", "sqlite"),
						huh.NewOption("MongoDB", "mongodb"),
					),
			),
		).
			WithAccessible(false).
			WithTheme(asset.SetTheme())

		fmt.Println(asset.VgoLogo)

		err := form.Run()
		if err != nil {
			fmt.Println(asset.Text.Foreground(asset.Red).Render(" Hey! Why stopped? "))
			return
		}

		_ = spinner.
			New().
			Title("Scaffolding the project").
			Action(func() {
				err = p.ScaffoldProject()
				if err != nil {
					fmt.Println(asset.Text.Foreground(asset.Red).Render(" Error : Sorry! Failed to scaffold the project"))
					res := p.RevertScaffold()
					if res != nil {
						fmt.Println(asset.Text.Foreground(asset.Red).Render(" Error : Failed to revert the scaffold"))
					} else {
						fmt.Println(asset.Text.Foreground(asset.Red).Render(" Reverted the scaffold"))
					}
				} else {
					fmt.Println(asset.Text.Foreground(asset.Green).Render(" Project initialized successfully"))
				}
			}).
			Style(asset.Text).
			Accessible(false).
			Run()
	},
}
