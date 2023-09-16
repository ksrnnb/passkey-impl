package middleware

import (
	"github.com/ksrnnb/passkey-impl/repository"
	"github.com/labstack/echo/v4"
)

func RepositoryMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(repository.RepositoriesContextName, repository.Repositories{
				UserRepository: repository.NewUserRepository(),
			})
			return next(c)
		}
	}
}
