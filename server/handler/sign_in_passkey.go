package handler

import (
	"encoding/json"
	"errors"
	"fmt"
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

func SignInPasskey(c echo.Context) error {
	response, err := protocol.ParseCredentialRequestResponseBody(c.Request().Body)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println(response.ParsedCredential.ID)
	credRepo := repository.Repos.CredentialRepository
	cred, err := credRepo.FindById(response.ParsedCredential.ID)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	userRepo := repository.Repos.UserRepository
	user, err := userRepo.FindById(cred.UserId)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return ErrorJSON(c, http.StatusBadRequest, "user not found")
		}
		return ErrorJSON(c, http.StatusInternalServerError, err.Error())
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

	_, err = w.ValidateLogin(user, session, response)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	token, err := jwt.New(user.Id)
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "jwt generation error")
	}

	res := signInPasskeyResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}
