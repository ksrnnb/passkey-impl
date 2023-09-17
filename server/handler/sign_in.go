package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ksrnnb/passkey-impl/jwt"
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

type signInRequest struct {
	UserId   string
	Password string
}

type signInResponse struct {
	Token string `json:"token"`
}

func SignIn(c echo.Context) error {
	repo, ok := c.Get(repository.RepositoriesContextName).(repository.Repositories)
	if !ok {
		return ErrorJSON(c, http.StatusInternalServerError, "unexpected error")
	}

	var req signInRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, "invalid request")
	}

	user, err := repo.UserRepository.FindById(req.UserId)
	if err != nil {
		return ErrorJSON(c, http.StatusBadRequest, "user not found")
	}

	// in this sample app, we don't validate password
	// if user.Password != c.FormValue("password") {
	// 	return ErrorJSON(c, http.StatusBadRequest, "user not found")
	// }

	token, err := jwt.New(user.Id)
	if err != nil {
		return ErrorJSON(c, http.StatusInternalServerError, "jwt generation error")
	}

	res := signInResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}
