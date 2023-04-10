package route

import (
	"net/http"
	"os"

	"echo_golang/controllers"
	m "echo_golang/middleware"
	"echo_golang/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()
	m.Logger(e)
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(dbHost),
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, login here !, localhost:8000/login")
	})

	c := controllers.UserStruct{}
	e.POST("/login", c.LoginUserController)
	e.GET("/users", c.GetUserController, echojwt.WithConfig(config))
	e.GET("users/auth", c.GetUserController, echojwt.WithConfig(config))
	e.POST("/users", c.CreateUserController, echojwt.WithConfig(config))
	e.DELETE("/users/:id", c.DeleteUserController, echojwt.WithConfig(config))
	e.PUT("/users/:id", c.UpdateUserController, echojwt.WithConfig(config))

	return e
}
