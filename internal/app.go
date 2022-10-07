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

// NewAppConfig provides a context with the app's db connection
func newAppConfig() *appConfig {
	ctx := context.Background()
	key := database.DatabaseCtxKey
	db := database.ProvideDatabase()

	// sqlxRepo := adapter.ProvideSqlxRepositoryAdapter(&sql.DB{})

	return &appConfig{
		Context:    context.WithValue(ctx, key, db),
		httpServer: restful.ProvideEchoServer(),
	}
}

func Start() {
	config := newAppConfig()
	config.httpServer.Start(":8080")
}
