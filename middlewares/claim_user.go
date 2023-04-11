package middleware

import (
	"echo_golang/models"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (*models.JwtCustomClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	token, _ := jwt.ParseWithClaims(tokenString, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(dbHost), nil
	})
	claims, _ := token.Claims.(*models.JwtCustomClaims)

	return claims, nil
}

func Restricted(c echo.Context) (string, error) {
	defer fmt.Println("end user")
	user := c.Get("user").(*jwt.Token)
	fmt.Println("failed user")
	claims := user.Claims.(*models.JwtCustomClaims)
	name := claims.Name
	return name, nil
}
