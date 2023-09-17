package model

type Passkey struct {
	Id           string
	UserId       string
	CredentialId string
	PublicKey    string
	SignInCount  int64
}
