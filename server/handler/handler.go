package handler

import (
	"github.com/labstack/echo/v4"
)

type errorMessage struct {
	Message string
}

func ErrorJSON(c echo.Context, httpStatusCode int, message string) error {
	msg := errorMessage{
		Message: message,
	}
	return c.JSON(httpStatusCode, msg)
}
