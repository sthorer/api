package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/config"
)

func JWTAuth(conf *config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ContextKey: types.JWTKey,
		SigningKey: []byte(conf.Secret),
		Claims:     &types.JWTCustomClaims{},
	})
}
