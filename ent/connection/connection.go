package connection

import (
	"context"
	"fmt"
	"log"

	"backend/ent"
	"backend/internal"
)

func NewDBConnection(env *internal.Environ) *ent.Client {
	// Create the DB connection options
	db, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBName, env.DBPass,
	))
	if err != nil {
		log.Printf("%v", err)
		log.Fatal("could not establish a DB connection in a timely manner")
	}

	// Run the auto migration
	if err := db.Schema.Create(context.Background()); err != nil {
		_ = db.Close()
		log.Printf("%v", err)
		log.Fatal("failed creating schema resources")
	}

	return db
}
