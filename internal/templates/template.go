package tmpl

import (
	_ "embed"
)

var (
	//go:embed app/app.go.tmpl
	AppTmpl string

	//go:embed cmd/server/server.go.tmpl
	ServerTmpl string

	//go:embed database/mongo.go.tmpl
	MongoTmpl string

	//go:embed database/mysql.go.tmpl
	MySqlTmpl string

	//go:embed database/postgres.go.tmpl
	PostgresTmpl string

	//go:embed database/sqlite.go.tmpl
	SqliteTmpl string

	//go:embed framework/chi.go.tmpl
	ChiTmpl string

	//go:embed framework/echo.go.tmpl
	EchoTmpl string

	////go:embed framework/fiber.go.tmpl
	FiberTmpl string

	//go:embed framework/gin.go.tmpl
	GinTmpl string
)
