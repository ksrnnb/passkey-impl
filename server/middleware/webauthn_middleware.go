package middleware

import (
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/handler"
	"github.com/labstack/echo/v4"
)

const RPOrigin = "http://localhost:3000"

func WebAuthnMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			wconfig := &webauthn.Config{
				RPDisplayName: "passkey-impl",     // Display Name for your site
				RPID:          "localhost",        // Generally the FQDN for your site
				RPOrigins:     []string{RPOrigin}, // The origin URLs allowed for WebAuthn requests
			}

			w, err := webauthn.New(wconfig)
			if err != nil {
				return handler.ErrorJSON(c, http.StatusInternalServerError, "webauthn middleware error")
			}

			c.Set(handler.WebAuthnContextKeyName, w)
			return next(c)
		}
	}
}
