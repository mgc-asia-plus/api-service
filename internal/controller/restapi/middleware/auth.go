package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	// HeaderAuthToken is the header key for simple auth (e.g. X-Auth-Token).
	HeaderAuthToken = "X-Auth-Token"
	// HeaderAuthorization is used when client sends "Bearer <token>".
	HeaderAuthorization = "Authorization"
	BearerPrefix         = "Bearer "
)

// Auth validates the request using a token sent in header.
// If token is empty, auth is skipped (all requests pass).
// Otherwise requires either:
//   - X-Auth-Token: <token>
//   - Authorization: Bearer <token>
func Auth(token string) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if token == "" {
			return ctx.Next()
		}

		got := ctx.Get(HeaderAuthToken)
		if got == "" {
			got = parseBearer(ctx.Get(HeaderAuthorization))
		}
		if got != token {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}
		return ctx.Next()
	}
}

func parseBearer(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, BearerPrefix) {
		return strings.TrimSpace(s[len(BearerPrefix):])
	}
	return ""
}
