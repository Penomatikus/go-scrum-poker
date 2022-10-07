package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed static
var staticDist embed.FS

func staticSource() http.Handler {
	fsys, err := fs.Sub(staticDist, "static")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(fsys))
}

func Mount(e *echo.Echo) {
	e.GET("/swagger-ui/*",
		echo.WrapHandler(http.StripPrefix("/swagger-ui/", staticSource())))
	e.GET("/swagger-ui", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/swagger-ui/")
	})
}
