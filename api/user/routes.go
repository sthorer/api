package user

import (
	"github.com/sthorer/api/api/middlewares"
	"github.com/sthorer/api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Apply(e *echo.Echo, conf *config.Config) {
	group := e.Group("/user")

	group.Use(middleware.JWT(conf.Secret))
	group.Use(middlewares.Auth)

	group.GET("/me", Me)
	group.POST("/token/new", NewToken)
	group.DELETE("/token/:id", RevokeToken)
}
