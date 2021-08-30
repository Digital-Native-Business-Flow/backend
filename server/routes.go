package server

import (
	"backend/server/handler"

	"github.com/gofiber/fiber/v2"
)

func assignRoutesAndHandlers(app *fiber.App, h *handler.Handler) {
	// Index
	app.Get("/", h.Index)

	// File upload
	app.Post("/upload", h.Upload)
}
