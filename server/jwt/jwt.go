package jwt

import (
	"errors"
	"time"

	jwtPkg "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var signingKey = []byte("sampleSigningKey")

const issuer = "passkey-impl"

const UserIdKey = "user_id"
const TokenKey = "token"

type Token struct {
	UserId string
}

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

func Parse(tokenString string) (Token, error) {
	token, err := jwtPkg.Parse(tokenString, func(token *jwtPkg.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return Token{}, err
	}
	if !token.Valid {
		return Token{}, errors.New("invalid token")
	}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		return Token{}, err
	}

	if isInvalidated(tokenString) {
		return Token{}, errors.New("invalid token")
	}

	return Token{UserId: subject}, nil
}

func newID() string {
	return uuid.New().String()
}
