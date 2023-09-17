package handler

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/labstack/echo/v4"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(stringLength int) string {
	b := make([]rune, stringLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateChallenge() string {
	return generateRandomString(32)
}

func challengeKvsKey(userId string) string {
	return fmt.Sprintf("user:%s:challenge", userId)
}

type ChallengeRegistrationResponse struct {
	UserId    string `json:"userId"`
	Challenge string `json:"challenge"`
}

func ChallengeRegistration(c echo.Context) error {
	userId, ok := c.Get(jwt.UserIdKey).(string)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}
	// repo, ok := c.Get(repository.RepositoriesContextName).(repository.Repositories)
	// if !ok {
	// 	return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	// }
	res := ChallengeRegistrationResponse{
		UserId:    userId,
		Challenge: generateChallenge(),
	}
	return c.JSON(http.StatusOK, res)
}

// TODO: implement
type RegisterPasskeyRequest struct {
}

func RegisterPasskey(c echo.Context) error {
	fmt.Println(c.Request().Body)

	return c.JSON(http.StatusOK, nil)
}
