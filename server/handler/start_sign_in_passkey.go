package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/kvs"
	"github.com/labstack/echo/v4"
)

type startSignInPasskeyRequest struct {
	UserId   string
	Password string
}

type startSignInPasskeyResponse struct {
	Token string `json:"token"`
}

func StartSignInPasskey(c echo.Context) error {
	w, ok := c.Get(WebAuthnContextKeyName).(*webauthn.WebAuthn)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	options, session, err := w.BeginDiscoverableLogin()
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	s, err := json.Marshal(session)
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}
	// store webauthn session to kvs
	kvs.Add(sessionKvsKey(session.Challenge), string(s))

	return c.JSON(http.StatusOK, options)
}
