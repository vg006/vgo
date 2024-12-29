package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	app "github.com/vg006/vgo/internal"
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
					Placeholder("Enter the project name").
					Description("Names the new project").
					Validate(utils.CheckValidProjectName),
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
				huh.NewNote().DescriptionFunc(func() string {
					switch p.FrameWork {
					case "stdlib":
						return "The standard library for building web applications in Go"
					case "echo":
						return "A high performance, extensible, minimalist web framework for Go"
					case "fiber":
						return "An Express inspired web framework for Go"
					case "chi":
						return "A lightweight, idiomatic and composable router for building Go HTTP services"
					case "gin":
						return "A web framework written in Go (Golang)"
					default:
						return ""
					}
				}, &p.FrameWork),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&p.Database).
					Title("Database").
					Description("Select the database for the project").
					Options(
						huh.NewOption("None", "None"),
						huh.NewOption("PostgreSQL", "PostgreSQL"),
						huh.NewOption("MySQL", "MySQL"),
						huh.NewOption("SQLite", "SQLite"),
						huh.NewOption("MongoDB", "MongoDB"),
					),
				huh.NewNote().DescriptionFunc(func() string {
					switch p.Database {
					case "None":
						return "No database will be used in the project/ Not decided yet"
					case "PostgreSQL":
						return "The world's most advanced open source database"
					case "MySQL":
						return "The world's most popular open source database"
					case "SQLite":
						return "A C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine"
					case "MongoDB":
						return "A general purpose, document-based, distributed database built for modern application developers and for the cloud era"
					default:
						return ""
					}
				}, &p.Database),
			),
		).WithAccessible(false)

		err := form.Run()
		if err != nil {
			fmt.Printf(" Error : Failed to initialize the project\n%s", err.Error())
		}

		if err = p.ScaffoldProject(); err != nil {
			fmt.Printf(" Error : Failed to scaffold the project\n%s", err.Error())
		}
	},
}
