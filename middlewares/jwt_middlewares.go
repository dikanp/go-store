package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GetSecretKey() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("JWT_KEY")
}

func GenerateToken(userId int, userEmail string) (string, error) {
	mapClaim := jwt.MapClaims{}
	mapClaim["userId"] = userId
	mapClaim["userEmail"] = userEmail
	mapClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)

	// key

	return token.SignedString([]byte(GetSecretKey()))
}
