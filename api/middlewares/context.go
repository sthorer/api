package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/config"
)

func Context(conf *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &types.Context{Context: c, Config: conf}
			return next(cc)
		}
	}
}
