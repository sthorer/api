package user

import (
	"github.com/sthorer/api/api/middlewares"
	"github.com/sthorer/api/config"

	"github.com/labstack/echo/v4"
)

func Apply(e *echo.Echo, conf *config.Config) {
	group := e.Group("/user")

	group.Use(middlewares.JWTAuth(conf))
	group.Use(middlewares.Auth)

	group.GET("/me", Me)
	group.GET("/tokens", ListTokens)
	group.POST("/tokens/new", NewToken)
	group.GET("/tokens/:id", GetToken)
	group.DELETE("/tokens/:id", RevokeToken)
	group.POST("/tokens/:id/reset", ResetToken)
}
