package model

import "github.com/go-webauthn/webauthn/webauthn"

type Credential struct {
	webauthn.Credential
	Name string
}
