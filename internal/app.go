package app

import (
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	tmpl "github.com/vg006/vgo/internal/templates"
)

func (p *Project) ScaffoldProject() error {
	// Creates the project root directory
	err := os.Mkdir(p.Name, 0754)
	if err != nil {
		return err
	}
	// -----------------------------------------------------------------
	// Changes the directory to the project root
	err = os.Chdir(p.Name)
	if err != nil {
		return err
	}
	// Initiates the project module
	_, err = exec.Command("go", "mod", "init", p.ModName).Output()
	if err != nil {
		return err
	}
	// Creates a Readme.md File
	f, err := os.Create("Readme.md")
	if err != nil {
		return err
	}
	defer f.Close()

	// -----------------------------------------------------------------
	// Creates the cmd directory
	go p.CreateCmdDir()
	// Creates the internal directory
	go p.CreateInternalDir()

	return nil
}

func (p *Project) CreateCmdDir() {
	// Creates the cmd directory
	err := os.Mkdir("cmd", 0754)
	if err != nil {
		panic(err)
	}

	// Creates the server directory
	serverPath := filepath.Join("cmd", "server")
	err = os.MkdirAll(serverPath, 0754)
	if err != nil {
		panic(err)
	}

	// Creates the server.go
	serverFile, err := os.Create(filepath.Join(serverPath, "server.go"))
	if err != nil {
		panic(err)
	}
	defer serverFile.Close()

	// Writes into serverFile
	err = template.
		Must(
			template.
				New("server.go").
				Funcs(tmpl.Functions).
				Parse(tmpl.ServerTmpl)).
		Execute(serverFile, p)
	if err != nil {
		panic(err)
	}
}

func (p *Project) CreateInternalDir() {
	// internal
	// -----------------------------------------------------------------
	// Creates the internal directory
	err := os.Mkdir("internal", 0754)
	if err != nil {
		panic(err)
	}

	// internal/app
	// -----------------------------------------------------------------
	// Creates the app directory
	appPath := filepath.Join("internal", "app")
	err = os.MkdirAll(appPath, 0754)
	if err != nil {
		panic(err)
	}
	// Creates the app.go
	appFile, err := os.Create(filepath.Join(appPath, "app.go"))
	if err != nil {
		panic(err)
	}
	defer appFile.Close()
	// Writes into appFile
	err = template.
		Must(
			template.
				New("app.go").
				Funcs(tmpl.Functions).
				Parse(tmpl.AppTmpl)).
		Execute(appFile, p)
	if err != nil {
		panic(err)
	}

	// internal/database
	// -----------------------------------------------------------------
	// Creates the database directory
	dbPath := filepath.Join("internal", "database")
	err = os.MkdirAll(dbPath, 0754)
	if err != nil {
		panic(err)
	}
	// Creates the database.go
	dbFile, err := os.Create(filepath.Join(dbPath, "database.go"))
	if err != nil {
		panic(err)
	}
	defer dbFile.Close()
	// Writes into dbFile
	err = template.
		Must(
			template.
				New("database.go").
				Funcs(tmpl.Functions).
				Parse(tmpl.DatabaseTmpl(p.Database))).
		Execute(dbFile, p)
	if err != nil {
		panic(err)
	}

	// internal/handlers
	// -----------------------------------------------------------------
	// Creates the handlers directory
	handlersPath := filepath.Join("internal", "handlers")
	err = os.MkdirAll(handlersPath, 0754)
	if err != nil {
		panic(err)
	}
	// Creates the handlers.go
	handlersFile, err := os.Create(filepath.Join(handlersPath, "handlers.go"))
	if err != nil {
		panic(err)
	}
	defer handlersFile.Close()
	// Writes into handlersFile
	err = template.
		Must(
			template.
				New("handlers.go").
				Funcs(tmpl.Functions).
				Parse(tmpl.HandlerTmpl(p.FrameWork))).
		Execute(handlersFile, p)
	if err != nil {
		panic(err)
	}
	// -----------------------------------------------------------------
}
