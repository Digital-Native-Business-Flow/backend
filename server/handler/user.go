package handler

import (
	"strconv"
	"time"

	"backend/database/model"
	"backend/internal"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login a user into the system and return an authorization JWT
func (h *Handler) Login(c *fiber.Ctx) (err error) {
	// Bind request data
	u := new(model.User)
	if err = c.BodyParser(&u); err != nil {
		return
	}

	// Validate request data
	if err = h.Validator.Validate(u); err != nil {
		return
	}

	// Find the user in the DB
	if err = model.UserFind(h.DB, u); err != nil {
		return
	}

	// JWT

	// Set claims
	claims := &internal.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.FormatInt(u.Id, 10),
			ExpiresAt: time.Now().Add(h.JwtExp).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	Token, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		return
	}

	return HTTPSuccess(c, fiber.Map{
		"Token": Token,
	})
}
