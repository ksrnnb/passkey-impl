package middleware

import (
	"github.com/labstack/echo/v4"
)

const TokenHeaderName = "Authorization"

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Get(TokenHeaderName)
			return next(c)
		}
	}
}
