package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

// NOTE:
const RelyingPartyName = "passkey-impl"
const RelyingPartyID = "localhost:8888"
const WebAuthnContextKeyName = "webauthn"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func challengeKvsKey(userId string) string {
	return fmt.Sprintf("user:%s:challenge", userId)
}

type ChallengeRegistrationResponse struct {
	UserId    string `json:"userId"`
	Challenge string `json:"challenge"`
}

func ChallengeRegistration(c echo.Context) error {
	user, err := CurrentUser(c)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return ErrorJSON(c, http.StatusBadRequest, "user not found")
		}
		return ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	w, ok := c.Get(WebAuthnContextKeyName).(*webauthn.WebAuthn)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	options, _, err := w.BeginRegistration(user)
	if err != nil {
		fmt.Println(err)
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	return c.JSON(http.StatusOK, options)
}

// TODO: implement
type RegisterPasskeyRequest struct {
}

func RegisterPasskey(c echo.Context) error {
	fmt.Println(c.Request().Body)

	return c.JSON(http.StatusOK, nil)
}
