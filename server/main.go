package main

import (
	"net/http"

	"github.com/ksrnnb/passkey-impl/handler"
	"github.com/ksrnnb/passkey-impl/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.RepositoryMiddleware())
	e.Use(middleware.AuthMiddleware())

	// Routes
	e.POST("/signin", handler.SignIn)
	e.POST("/signout", hello)
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
