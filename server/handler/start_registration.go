package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/kvs"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

const WebAuthnContextKeyName = "webauthn"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func sessionKvsKey(challenge string) string {
	return fmt.Sprintf("challenge:%s", challenge)
}

func StartRegistration(c echo.Context) error {
	user, err := CurrentUser(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return ErrorJSON(c, http.StatusBadRequest, "user not found")
		}
		return ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// TODO: add exclusions
	// credRepo := repository.Repos.CredentialRepository
	// creds := credRepo.FindByUserId(user.Id)
	// webauthn.WithExclusions(creds)

	w, ok := c.Get(WebAuthnContextKeyName).(*webauthn.WebAuthn)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	options, session, err := w.BeginRegistration(user)
	if err != nil {
		fmt.Println(err)
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	s, err := json.Marshal(session)
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	// store webauthn session to kvs
	kvs.Add(sessionKvsKey(session.Challenge), string(s))

	return c.JSON(http.StatusOK, options)
}
