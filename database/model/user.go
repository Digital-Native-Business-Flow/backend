package model

import (
	"backend/internal"
	"backend/security"

	"github.com/go-pg/pg/v10"
)

type User struct {
	tableName struct{} `pg:"users,alias:users"`
	Id        int64
	Email     string `validate:"required,email"`
	Password  string `validate:"required,password"`
	Salt      []byte
}

// Find a user in the DB by email and password
func UserFind(db *pg.DB, u *User) error {
	// Save the raw password for later use
	rawPassword := u.Password

	// Search the DB for the user using email
	err := db.Model(u).Where("users.email = ?", u.Email).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return internal.NewError(internal.ErrBEInvalidPassword, err, 1)
		}

		return internal.NewError(internal.ErrDBQuery, err, 1)
	}

	// Create the hashed password for the current user
	hashedPassword := security.HashPassword(rawPassword, u.Salt)

	// Check if hashed passwords match
	if u.Password != hashedPassword {
		return internal.NewError(internal.ErrBEInvalidPassword, nil, 1)
	}

	return nil
}

// Create a new root user in the DB
func UserCreateRoot(db *pg.DB, u *User) error {
	// Add the root user to the DB
	_, err := db.Model(u).Insert()
	if err != nil {
		return internal.NewError(internal.ErrDBInsert, err, 1)
	}

	return nil
}
