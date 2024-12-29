package template

import (
	_ "embed"
)

var (
	//go:embed app/app.go.tmpl
	AppTmpl []byte

	//go:embed cmd/server/server.go.tmpl
	ServerTmpl []byte

	//go:embed database/mongo.go.tmpl
	MongoTmpl []byte

	//go:embed database/mysql.go.tmpl
	MySqlTmpl []byte

	//go:embed database/postgres.go.tmpl
	PostgresTmpl []byte

	//go:embed database/sqlite.go.tmpl
	SqliteTmpl []byte

	//go:embed framework/chi.go.tmpl
	ChiTmpl []byte

	//go:embed framework/echo.go.tmpl
	EchoTmpl []byte

	////go:embed framework/fiber.go.tmpl
	FiberTmpl []byte

	//go:embed framework/gin.go.tmpl
	GinTmpl []byte
)
