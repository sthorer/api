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
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["name"].(int)

		u, err := c.Client.User.Get(context.Background(), id)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		c.Set(types.UserKey, u)
		return next(c)
	}
}
