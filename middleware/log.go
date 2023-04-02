package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func TestJWT(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("lewat middleware")
	return next
}
