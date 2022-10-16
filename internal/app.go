package internal

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/penomatikus/go-scrum-poker/server/database"
	"github.com/penomatikus/go-scrum-poker/server/restful"
)

type appConfig struct {
	Context    context.Context
	httpServer *echo.Echo
}

// NewAppConfig provides a context with the app's db connection. When no DB connection could be established
// the app will panic and die.
func newAppConfig() *appConfig {
	//TODO: make this configurable
	dataSourceName := "scrumpoker:password@tcp(localhost:3306)/database?charset=utf8&parseTime=True&loc=Local"

	ctx := context.Background()
	key := database.DatabaseCtxKey
	orm, err := database.ProvideORM(database.SQLX, dataSourceName)

	if err != nil {
		panic(err)
	}

	return &appConfig{
		Context:    context.WithValue(ctx, key, orm),
		httpServer: restful.ProvideEchoServer(),
	}
}

func Start() {
	config := newAppConfig()
	config.httpServer.Start(":8080")
}
