package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/ksrnnb/passkey-impl/kvs"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

type signInPasskeyRequest struct {
	UserId   string
	Password string
}

type signInPasskeyResponse struct {
	Token string `json:"token"`
}

func discoverableUserHandler(rawID, userHandle []byte) (user webauthn.User, err error) {
	userRepo := repository.Repos.UserRepository
	return userRepo.FindById(string(userHandle))
}

func SignInPasskey(c echo.Context) error {
	response, err := protocol.ParseCredentialRequestResponseBody(c.Request().Body)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	sessionString, err := kvs.Get(sessionKvsKey(response.Response.CollectedClientData.Challenge))
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "session not found")
	}

	var session webauthn.SessionData
	if err := json.Unmarshal([]byte(sessionString), &session); err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "json parse error")
	}

	w, ok := c.Get(WebAuthnContextKeyName).(*webauthn.WebAuthn)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	_, err = w.ValidateDiscoverableLogin(discoverableUserHandler, session, response)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	token, err := jwt.New(string(response.Response.UserHandle))
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "jwt generation error")
	}

	res := signInPasskeyResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}
