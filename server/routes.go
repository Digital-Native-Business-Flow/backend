package server

import (
	"backend/server/handler"

	"github.com/gofiber/fiber/v2"
)

func assignRoutesAndHandlers(app *fiber.App, h *handler.Handler) {
	// Index
	app.Get("/", h.Index)

	// Users processes management
	app.Post("/login", h.Login)

	// Restricted
	app.Get("/restricted", h.Restricted)
}
