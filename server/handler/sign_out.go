package handler

import (
	"net/http"

	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/labstack/echo/v4"
)

func SignOut(c echo.Context) error {
	tokenString, ok := c.Get(jwt.TokenKey).(string)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	jwt.Invalidate(tokenString)

	return c.JSON(http.StatusOK, nil)
}
