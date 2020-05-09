package files

import (
	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/api/middlewares"
)

func Apply(e *echo.Echo) {
	group := e.Group("/files")

	group.Use(middlewares.TokenAuth())

	group.POST("/upload", Upload)
}
