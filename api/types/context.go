package types

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/config"
)

const (
	UserKey  = "user"
	TokenKey = "token"
)

type Context struct {
	echo.Context
	*config.Config
}

func (c *Context) ValidationError(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
