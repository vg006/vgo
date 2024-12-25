package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Project struct {
	Name      string
	FrameWork string
	DataBase  string
	Protocol  string
	Addons    []string
}

func NewProject(name, framework, database string, addons []string) *Project {
	return &Project{
		Name:      name,
		FrameWork: framework,
		DataBase:  database,
		Addons:    addons,
	}
}

func Init() {
	var p Project

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&p.Name).
				Title("Project Name").
				Placeholder("Enter the project name").
				Description("Names the new project").
				Validate(func(s string) error {
					if s == "" {
						return errors.New("No null name, sorry")
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
				Value(&p.DataBase).
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
				switch p.DataBase {
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
			}, &p.DataBase),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Value(&p.Addons).
				Title("Addons").
				Description("Select the addons for the project").
				Options(
					huh.NewOption("None", "None"),
					huh.NewOption("Auth", "Auth"),
					huh.NewOption("Email", "Email"),
					huh.NewOption("Caching", "Caching"),
					huh.NewOption("Logging", "Logging"),
					huh.NewOption("Testing", "Testing"),
				),
			huh.NewNote().DescriptionFunc(func() string {
				switch len(p.Addons) {
				case 0:
					return "No addons will be used in the project/ Not decided yet"
				case 1:
					return "The following addon will be used in the project"
				default:
					return "The following addons will be used in the project"
				}
			}, &p.Addons),
		),
	).WithAccessible(false)

	err := form.Run()
	if err != nil {
		fmt.Println(" Error : Failed to initialize the project", err)
	}

	scaffoldProject := func() {
		// Simulate some work
		time.Sleep(4 * time.Second)
	}

	_ = spinner.New().Title("Preparing the project").Action(scaffoldProject).Run()

}
