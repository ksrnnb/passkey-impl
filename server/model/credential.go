package model

import (
	"encoding/base64"

	"github.com/go-webauthn/webauthn/webauthn"
)

type Credential struct {
	webauthn.Credential
	UserId string
	Name   string
}

func (c Credential) Id() string {
	return base64.RawURLEncoding.EncodeToString(c.ID)
}
