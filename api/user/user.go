package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/ent"
)

func Me(ctx echo.Context) error {
	c := ctx.(*types.Context)
	return c.JSON(http.StatusOK, c.Get(types.UserKey).(*ent.User))
}

func ListTokens(ctx echo.Context) error {
	c := ctx.(*types.Context)
	u := c.Get(types.UserKey).(*ent.User)

	tokens, err := c.Client.Token.
		Query().
		Where(token.HasUserWith(user.ID(u.ID))).
		All(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tokens)
}

func NewToken(ctx echo.Context) error {
	c := ctx.(*types.Context)
	user := c.Get(types.UserKey).(*ent.User)

	var body types.NewTokenRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	fmt.Println("USER:", user)

	token, err := c.Client.NewToken(context.Background(), user, body.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &types.TokenSecretResponse{
		Token:  token,
		Secret: token.Secret,
	})
}

func ResetToken(c echo.Context) error {
	cc := c.(*types.Context)
	id, err := uuid.FromBytes([]byte(cc.Param("id")))
	if err != nil {
		return cc.NoContent(http.StatusNotFound)
	}

	ctx := context.Background()
	u := cc.Get(types.UserKey).(*ent.User)
	t, err := cc.Client.Token.
		Query().
		Where(token.ID(id), token.HasUserWith(user.ID(u.ID))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return cc.NoContent(http.StatusNotFound)
		}
		return err
	}

	t, err = cc.Client.ResetToken(ctx, t)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &types.TokenSecretResponse{
		Token:  t,
		Secret: t.Secret,
	})
}

func GetToken(ctx echo.Context) error {
	c := ctx.(*types.Context)
	id, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	u := c.Get(types.UserKey).(*ent.User)
	t, err := c.Client.Token.
		Query().
		Where(token.ID(id), token.HasUserWith(user.ID(u.ID))).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		return err
	}

	return c.JSON(http.StatusOK, t)
}

func RevokeToken(ctx echo.Context) error {
	c := ctx.(*types.Context)
	id, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Client.RevokeToken(context.Background(), id); err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return err
	}

	return c.NoContent(http.StatusOK)
}
