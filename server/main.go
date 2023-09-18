package main

import (
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
	e.Use(echoMiddleware.CORS())
	e.Use(middleware.WebAuthnMiddleware())

	// Unauthenticated Routes
	e.POST("/signin", handler.SignIn)

	// Authenticated Routes
	e.POST("/authenticated", handler.AuthenticatedUser, middleware.AuthMiddleware())
	e.POST("/signout", handler.SignOut, middleware.AuthMiddleware())
	e.POST("/passkey/register/start", handler.StartRegistration, middleware.AuthMiddleware())
	e.POST("/passkey/register", handler.RegisterPasskey, middleware.AuthMiddleware())
	e.DELETE("/passkey/:id", handler.DeletePasskey, middleware.AuthMiddleware())

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
