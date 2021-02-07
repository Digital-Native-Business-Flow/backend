package handler

import (
	"time"

	"backend/internal"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	Handler struct {
		JwtSecret string
		JwtExp    time.Duration
		DB        *pg.DB
		Validator *internal.Validator
	}
)

// Return a formatted HTTP Success response
func HTTPSuccess(c *fiber.Ctx, d interface{}) error {
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Action completed successfully",
		"data":    d,
	})
}
