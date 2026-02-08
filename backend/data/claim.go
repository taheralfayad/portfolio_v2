package data

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

