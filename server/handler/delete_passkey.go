package handler

import (
	"net/http"

	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

func DeletePasskey(c echo.Context) error {
	credId := c.Param("id")
	credRepo := repository.Repos.CredentialRepository
	credRepo.Delete(credId)

	return c.JSON(http.StatusOK, nil)
}
