package model

import "github.com/dgrijalva/jwt-go"

// Jwt json web token
type Jwt struct {
	ID   int
	Name string
	jwt.StandardClaims
}
