package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/kvs"
	"github.com/ksrnnb/passkey-impl/model"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

func RegisterPasskey(c echo.Context) error {
	response, err := protocol.ParseCredentialCreationResponseBody(c.Request().Body)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

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

	sessionString, err := kvs.Get(sessionKvsKey(user.Id))
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "session not found")
	}

	var session webauthn.SessionData
	if err := json.Unmarshal([]byte(sessionString), &session); err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "json parse error")
	}

	credential, err := w.CreateCredential(user, session, response)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	// update credential
	user.Credentials = append(user.Credentials, model.Credential{
		Credential: *credential,
		Name:       c.Request().UserAgent(),
	})

	userRepo := repository.Repos.UserRepository
	userRepo.Add(user)

	return c.JSON(http.StatusOK, credential)
}
