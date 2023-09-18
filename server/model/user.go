package model

import "github.com/go-webauthn/webauthn/webauthn"

type User struct {
	Id          string
	Name        string
	Password    string
	Credentials []Credential
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.Id)
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.Name
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	creds := make([]webauthn.Credential, len(u.Credentials))
	for i, cred := range u.Credentials {
		creds[i] = cred.Credential
	}
	return creds
}
