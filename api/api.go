package api

import (
	"net/http"

	"github.com/sthorer/api/api/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sthorer/api/api/auth"
	"github.com/sthorer/api/api/files"
	"github.com/sthorer/api/api/middlewares"
	"github.com/sthorer/api/api/types"

	"github.com/sthorer/api/config"
)

func New(conf *config.Config) *echo.Echo {
	e := echo.New()

	e.Validator = &types.Validator{Validator: conf.Validator}
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middlewares.Context(conf))

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	auth.Apply(e)
	user.Apply(e, conf)
	files.Apply(e)

	return e
}
