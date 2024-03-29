package security

import (
	"backend/internal"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func ConfigureSecurity(app *fiber.App, env *internal.Environ) {
	app.Use(jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			// Skip authentication for certain request routes
			return internal.StringInSlice(c.Path(), []string{"/", "/login"})
		},
		SigningKey:    []byte(env.JwtSecret),
		SigningMethod: "HS512",
		Claims:        &internal.JWTClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return internal.NewError(internal.ErrBEJwtBad, err, 1)
			} else {
				return internal.NewError(internal.ErrBEJwtInvalid, err, 1)
			}
		},
	}))
}
