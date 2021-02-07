package database

import (
	"context"
	"crypto/tls"
	"log"

	"backend/internal"

	"github.com/go-pg/pg/v10"
)

func NewDBConnection(ctx context.Context, env *internal.Environ) *pg.DB {
	// Create the DB connection options
	opts := &pg.Options{
		Addr:      env.DBAddr,
		User:      env.DBUser,
		Password:  env.DBPass,
		Database:  env.DBName,
		TLSConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
	}

	// Connect to the DB
	db := pg.Connect(opts)

	// Check that the DB connection has been established
	err := db.Ping(ctx)
	if err != nil {
		log.Printf("%v", err)
		log.Fatal("could not establish a DB connection in a timely manner")
	}

	return db
}
