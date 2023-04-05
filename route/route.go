package route

import (
	"net/http"
	"os"

	"echo_golang/controllers"
	m "echo_golang/middleware"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()
	m.Logger(e)
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, login here !, localhost:8000/login")
	})

	e.POST("/login", controllers.LoginUserController)
	e.GET("/users", controllers.GetUserController, echojwt.JWT([]byte(dbHost)))
	e.GET("users/auth", controllers.GetUserController, echojwt.JWT([]byte(dbHost)))
	e.POST("/users", controllers.CreateUserController, echojwt.JWT([]byte(dbHost)))
	e.DELETE("/users/:id", controllers.DeleteUserController, echojwt.JWT([]byte(dbHost)))
	e.PUT("/users/:id", controllers.UpdateUserController, echojwt.JWT([]byte(dbHost)))

	return e
}
