package restful

import (
	"github.com/penomatikus/go-scrum-poker/server/swaggerui"

	"github.com/labstack/echo/v4"
)

func ProvideEchoServer() *echo.Echo {
	echo := echo.New()
	swaggerui.Mount(echo)
	return echo
}
