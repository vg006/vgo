package app

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"text/template"

	tmpl "github.com/vg006/vgo/internal/templates"
)

var (
	errChan = make(chan error, 2)
)

func (p *Project) ScaffoldProject() error {
	if err := os.Mkdir(p.Name, 0754); err != nil {
		return err
	}

	if err := os.Chdir(p.Name); err != nil {
		return err
	}

	if _, err := exec.Command("go", "mod", "init", p.ModName).Output(); err != nil {
		return err
	}

	f, err := os.Create("README.md")
	if err != nil {
		return err
	}

	defer f.Close()

	err = template.
		Must(
			template.New("README.md").Parse(tmpl.ReadmeMdTmpl)).
		Execute(f, p)
	if err != nil {
		errChan <- err
	}

	f, err = os.Create(".env")
	if err != nil {
		return err
	}
	defer f.Close()

	err = template.
		Must(
			template.
				New(".env").
				Parse(tmpl.EnvTmpl)).
		Execute(f, p)
	if err != nil {
		errChan <- err
	}

	f, err = os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer f.Close()

	err = template.
		Must(
			template.
				New(".gitignore").
				Parse(tmpl.GitignoreTmpl)).
		Execute(f, p)
	if err != nil {
		errChan <- err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		p.CreateCmdDir()
	}()

	go func() {
		defer wg.Done()
		p.CreateInternalDir()
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	if _, err = exec.Command("go", "mod", "tidy").Output(); err != nil {
		return err
	}

	if _, err = exec.Command("go", "fmt", "./...").Output(); err != nil {
		return err
	}

	return nil
}

func (p *Project) CreateCmdDir() {
	err := os.Mkdir("cmd", 0754)
	if err != nil {
		errChan <- err
	}

	serverPath := filepath.Join("cmd", "server")
	err = os.MkdirAll(serverPath, 0754)
	if err != nil {
		errChan <- err
	}

	serverFile, err := os.Create(filepath.Join(serverPath, "server.go"))
	if err != nil {
		errChan <- err
	}
	defer serverFile.Close()

	err = template.
		Must(
			template.
				New("server.go").
				Parse(tmpl.ServerTmpl)).
		Execute(serverFile, p)
	if err != nil {
		errChan <- err
	}
}

func (p *Project) CreateInternalDir() {
	err := os.Mkdir("internal", 0754)
	if err != nil {
		errChan <- err
	}

	appPath := filepath.Join("internal", "app")
	err = os.MkdirAll(appPath, 0754)
	if err != nil {
		errChan <- err
	}

	appFile, err := os.Create(filepath.Join(appPath, "app.go"))
	if err != nil {
		errChan <- err
	}
	defer appFile.Close()

	err = template.
		Must(
			template.
				New("app.go").
				Parse(tmpl.AppTmpl)).
		Execute(appFile, p)
	if err != nil {
		errChan <- err
	}

	dbPath := filepath.Join("internal", "database")
	err = os.MkdirAll(dbPath, 0754)
	if err != nil {
		errChan <- err
	}

	dbFile, err := os.Create(filepath.Join(dbPath, "database.go"))
	if err != nil {
		errChan <- err
	}
	defer dbFile.Close()

	err = template.
		Must(
			template.
				New("database.go").
				Parse(tmpl.DatabaseTmpl(p.Database))).
		Execute(dbFile, p)
	if err != nil {
		errChan <- err
	}

	handlersPath := filepath.Join("internal", "handlers")
	err = os.MkdirAll(handlersPath, 0754)
	if err != nil {
		errChan <- err
	}

	handlersFile, err := os.Create(filepath.Join(handlersPath, "handlers.go"))
	if err != nil {
		errChan <- err
	}
	defer handlersFile.Close()

	err = template.
		Must(
			template.
				New("handlers.go").
				Parse(tmpl.HandlerTmpl(p.FrameWork))).
		Execute(handlersFile, p)
	if err != nil {
		errChan <- err
	}
}

func (p *Project) RevertScaffold() error {
	if err := os.Chdir(".."); err != nil {
		return err
	}

	if err := os.RemoveAll(p.Name); err != nil {
		return err
	}
	return nil
}
