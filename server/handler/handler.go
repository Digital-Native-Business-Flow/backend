package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/ent"
	"backend/internal"
)

type (
	Handler struct {
		JwtSecret string
		JwtExp    time.Duration
		DB        *ent.Client
		Validator *internal.Validator
	}
)

// HTTPSuccess returns a formatted HTTP Success response
func HTTPSuccess(c *fiber.Ctx, d interface{}) error {
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Action completed successfully",
		"data":    d,
	})
}
