package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/ent"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

func TokenAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, secret string, c echo.Context) (bool, error) {
		cc := c.(*types.Context)
		ctx := context.Background()
		t, err := cc.Client.Token.
			Query().
			Where(token.Secret(secret), token.HasUserWith(user.Email(username))).
			WithUser().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return false, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			return false, err
		}

		if err = t.Update().SetLastUsed(time.Now()).Exec(ctx); err != nil {
			return false, err
		}

		cc.Set(types.TokenKey, t)
		cc.Set(types.UserKey, t.Edges.User)

		return true, nil
	})
}
