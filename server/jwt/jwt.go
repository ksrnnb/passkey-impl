package jwt

import (
	"time"

	jwtPkg "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var signingKey = []byte("sampleSigningKey")

const issuer = "passkey-impl"

func New(subject string) (string, error) {
	now := time.Now()

	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, jwtPkg.RegisteredClaims{
		ID:        newID(),
		Issuer:    issuer,
		Subject:   subject,
		ExpiresAt: jwtPkg.NewNumericDate(now.Add(24 * time.Hour)),
		NotBefore: jwtPkg.NewNumericDate(now),
		IssuedAt:  jwtPkg.NewNumericDate(now),
	})

	return token.SignedString(signingKey)
}

func newID() string {
	return uuid.New().String()
}
