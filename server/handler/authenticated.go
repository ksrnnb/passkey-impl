package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authenticated(c echo.Context) error {
	// middleware checks whether jwt token is valid.
	// so handler just return nil value
	return c.JSON(http.StatusOK, nil)
}
