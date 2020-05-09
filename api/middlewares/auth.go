package middlewares

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/api/types"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := ctx.(*types.Context)
		user := c.Get(types.JWTKey).(*jwt.Token)
		claims := user.Claims.(*types.JWTCustomClaims)

		u, err := c.Client.User.Get(context.Background(), claims.UserID)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		c.Set(types.UserKey, u)

		return next(c)
	}
}
