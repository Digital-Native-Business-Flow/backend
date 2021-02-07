package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"backend/server"

	"github.com/sirupsen/logrus"
)

func main() {
	// Configure base logging
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	// Instantiate the Fiber REST API server and DB connection
	app, serverPort, dbConn := server.InitServer()

	// Configure graceful shutdown of the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		<-quit
		logrus.Infof("Gracefully shutting down the server")

		if err := app.Shutdown(); err != nil {
			logrus.Fatalf("Failed to gracefully shut down the server: %s", err)
		}

		if err := dbConn.Close(); err != nil {
			logrus.Fatalf("Failed to gracefully disconnect from the DB: %s", err)
		}
	}()

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%d", serverPort)); err != nil {
		logrus.Fatalf("Failed to start the server: %s", err)
	}

	logrus.Infof("Server shut down complete")
}
