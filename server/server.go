package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/ent"
	"backend/ent/connection"
	"backend/internal"
	"backend/server/handler"
)

// Initialize a Fiber server and DB connection
func InitServer() (*fiber.App, int, *ent.Client) {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:          internal.ErrorHandler,
		Prefork:               false,
		DisableStartupMessage: false,
		BodyLimit:             100 * 1024 * 1024,
	})

	// Get configuration from environment
	env, err := internal.GetEnv()
	if err != nil {
		logrus.Fatalf("Failed to get environment variables: %s", err)
	}

	// Configure Fiber
	configureFiber(app)

	// Configure security
	//security.ConfigureSecurity(app, env)

	// Create a DB connection
	dbConn := connection.NewDBConnection(env)

	// Configure the request validator
	v, err := internal.NewValidator()
	if err != nil {
		logrus.Fatalf("Failed to create request validator: %s", err)
	}

	// Initialize the route handler
	h := &handler.Handler{
		JwtSecret: env.JwtSecret,
		JwtExp:    env.JwtExp,
		DB:        dbConn,
		Validator: v,
	}

	// Assign the routes & handlers
	assignRoutesAndHandlers(app, h)

	// Return
	return app, env.ServerPort, dbConn
}
