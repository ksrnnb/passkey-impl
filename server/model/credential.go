package model

import "github.com/go-webauthn/webauthn/webauthn"

type Credential struct {
	webauthn.Credential
	Id     string
	UserId string
	Name   string
}
