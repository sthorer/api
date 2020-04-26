package auth

import "github.com/labstack/echo/v4"

func Apply(e *echo.Echo) {
	group := e.Group("/auth")

	group.POST("/login", Login)
	group.POST("/register", Register)
}
