package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/ent"
)

func Me(ctx echo.Context) error {
	c := ctx.(*types.Context)
	return c.JSON(http.StatusOK, c.Get(types.UserKey).(*ent.User))
}

func NewToken(ctx echo.Context) error {
	c := ctx.(*types.Context)
	user := c.Get(types.UserKey).(*ent.User)

	token, err := c.Client.NewToken(context.Background(), user, "")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func RevokeToken(ctx echo.Context) error {
	c := ctx.(*types.Context)

	tokenID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if err := c.Client.RevokeToken(context.Background(), tokenID); err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return err
	}

	return c.NoContent(http.StatusOK)
}
