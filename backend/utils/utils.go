package utils

import (
	"os"
	"fmt"
	"path/filepath"
	"encoding/base64"
	"strings"
	"errors"
	"time"

	data "github.com/taheralfayad/portfolio_v2/data"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SaveBase64ImageToDisk(
	b64 string,
	path string,
) error {
	data, err := DecodeBase64Image(b64)
	if err != nil {
		return err
	}
	
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}

func DecodeBase64Image(b64 string) ([]byte, error) {
	if strings.Contains(b64, ",") {
		parts := strings.SplitN(b64, ",", 2)
		b64 = parts[1]
	}

	return base64.StdEncoding.DecodeString(b64)
}

func GenerateJWT(userID int, name string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	fmt.Println(jwtSecret)

	claims := data.Claims{
		UserID: userID,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    "portfolio",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateJWT(tokenString string) (*data.Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET not set")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&data.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*data.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
