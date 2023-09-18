package handler

import (
	"errors"

	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/ksrnnb/passkey-impl/model"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

type errorMessage struct {
	Message string
}

func ErrorJSON(c echo.Context, httpStatusCode int, message string) error {
	msg := errorMessage{
		Message: message,
	}
	return c.JSON(httpStatusCode, msg)
}

func CurrentUser(c echo.Context) (*model.User, error) {
	userId, ok := c.Get(jwt.UserIdKey).(string)
	if !ok {
		return nil, errors.New("cannot get user id from context")
	}

	repo, ok := c.Get(repository.RepositoriesContextName).(repository.Repositories)
	if !ok {
		return nil, errors.New("repository middleware error")
	}

	return repo.UserRepository.FindById(userId)
}
