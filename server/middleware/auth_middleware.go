package middleware

import (
	"net/http"
	"strings"

	"github.com/ksrnnb/passkey-impl/handler"
	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/labstack/echo/v4"
)

const tokenHeaderName = "Authorization"

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenWithBearer := c.Request().Header.Get(tokenHeaderName)
			if tokenWithBearer == "" {
				return handler.ErrorJSON(c, http.StatusBadRequest, "token not found")
			}

			tokenString := strings.TrimPrefix(tokenWithBearer, "Bearer ")
			token, err := jwt.Parse(tokenString)
			if err != nil {
				return handler.ErrorJSON(c, http.StatusBadRequest, "invalid token")
			}

			c.Set(jwt.UserIdKey, token.UserId)
			c.Set(jwt.TokenKey, tokenString)

			return next(c)
		}
	}
}
