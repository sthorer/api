package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/sthorer/api/api/types"
	"github.com/sthorer/api/utils"
)

func Login(ctx echo.Context) error {
	c := ctx.(*types.Context)

	var auth types.AuthRequest
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(&auth); err != nil {
		return c.ValidationError(err)
	}

	u, err := c.Client.UserLogin(context.Background(), auth.Email, auth.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid email or password")
	}

	token, err := utils.GenerateJWT(c.Config.Secret, u.ID, time.Now().Add(time.Hour*24*7))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &types.AuthResponse{User: u, Token: token})
}

func Register(ctx echo.Context) error {
	c := ctx.(*types.Context)

	var auth types.AuthRequest
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(&auth); err != nil {
		return c.ValidationError(err)
	}

	u, err := c.Client.UserRegister(context.Background(), auth.Email, auth.Password)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	token, err := utils.GenerateJWT(c.Config.Secret, u.ID, time.Now().Add(time.Hour*24*7))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &types.AuthResponse{User: u, Token: token})
}
