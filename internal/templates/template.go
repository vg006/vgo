package tmpl

import (
	_ "embed"
)

var (
	//go:embed app/app.go.tmpl
	AppTmpl string
	//go:embed cmd/server/server.go.tmpl
	ServerTmpl string
	//go:embed files/env.tmpl
	EnvTmpl string
	//go:embed files/readme.md.tmpl
	ReadmeMdTmpl string
	//go:embed files/gitignore.tmpl
	GitignoreTmpl string

	// Database Templates
	//go:embed database/mongo.go.tmpl
	MongoTmpl string
	//go:embed database/mysql.go.tmpl
	MySqlTmpl string
	//go:embed database/postgres.go.tmpl
	PostgresTmpl string
	//go:embed database/sqlite.go.tmpl
	SqliteTmpl string
	//go:embed database/none.go.tmpl
	NoneTmpl string

	// Handler Templates
	//go:embed framework/stdlib.go.tmpl
	StdLibTmpl string
	//go:embed framework/chi.go.tmpl
	ChiTmpl string
	//go:embed framework/echo.go.tmpl
	EchoTmpl string
	//go:embed framework/fiber.go.tmpl
	FiberTmpl string
	//go:embed framework/gin.go.tmpl
	GinTmpl string
)

func DatabaseTmpl(db string) string {
	switch db {
	case "mysql":
		return MySqlTmpl
	case "postgresql":
		return PostgresTmpl
	case "sqlite":
		return SqliteTmpl
	case "mongo":
		return MongoTmpl
	default:
		return NoneTmpl
	}
}

func HandlerTmpl(handler string) string {
	switch handler {
	case "echo":
		return EchoTmpl
	case "gin":
		return GinTmpl
	case "fiber":
		return FiberTmpl
	case "chi":
		return ChiTmpl
	case "stdlib":
		return StdLibTmpl
	default:
		return ""
	}
}
