package internal

import (
	"github.com/form3tech-oss/jwt-go"
)

type (
	JWTClaims struct {
		jwt.StandardClaims
	}
)
