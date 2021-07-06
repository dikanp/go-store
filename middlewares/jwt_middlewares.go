package middlewares

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func GetSecretKey() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to identify current directory (needed to load .env.test)")
		os.Exit(1)
	}
	basepath := filepath.Dir(file)
	err := godotenv.Load(filepath.Join(basepath, "../.env"))
	fmt.Println(err)
	if err != nil {
		// if err := godotenv.Load("../env"); err != nil {
			fmt.Println(err)
			log.Fatal("Error loading .env file")
		// }
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

func HandleLogin(c echo.Context) (int, string) {
	user := c.Get("customer").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		email := claims["email"].(string)
		return int(userId), email
	}
	return 0, ""

}
