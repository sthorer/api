package middlewares

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/ent"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

func File() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, tkn string, ctx echo.Context) (bool, error) {
		c := ctx.(*types.Context)
		t, err := c.Client.Token.
			Query().
			Where(token.Token(tkn), token.HasUserWith(user.Email(username))).
			Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return false, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			return false, err
		}

		c.Set(types.TokenKey, t)
		return true, nil
	})
}
