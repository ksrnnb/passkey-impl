package handler

import (
	"errors"
	"net/http"

	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

type PasskeyCredential struct {
	Name string `json:"name"`
}

type AuthenticatedUserResponse struct {
	UserId      string              `json:"userId"`
	Credentials []PasskeyCredential `json:"credentials"`
}

func AuthenticatedUser(c echo.Context) error {
	user, err := CurrentUser(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return ErrorJSON(c, http.StatusBadRequest, "user not found")
		}
		return ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	creds := make([]PasskeyCredential, len(user.Credentials))
	for i, cred := range user.Credentials {
		creds[i] = PasskeyCredential{
			Name: cred.Name,
		}
	}

	res := AuthenticatedUserResponse{
		UserId:      user.Id,
		Credentials: creds,
	}

	return c.JSON(http.StatusOK, res)
}
