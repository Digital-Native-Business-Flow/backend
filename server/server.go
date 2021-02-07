package server

import (
	"backend/security"
	"context"
	"time"

	"backend/database"
	"backend/internal"
	"backend/server/handler"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Initialize a Fiber server and DB connection
func InitServer() (*fiber.App, int, *pg.DB) {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:          internal.ErrorHandler,
		Prefork:               false,
		DisableStartupMessage: false,
	})

	// Get configuration from environment
	env, err := internal.GetEnv()
	if err != nil {
		logrus.Fatalf("Failed to get environment variables: %s", err)
	}

	// Configure Fiber
	configureFiber(app)

	// Configure security
	security.ConfigureSecurity(app, env)

	// Create a DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConn := database.NewDBConnection(ctx, env)

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
