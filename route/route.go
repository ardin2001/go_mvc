package route

import (
	"net/http"
	"os"

	"echo_golang/controllers"
	m "echo_golang/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
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
	e.GET("/users", controllers.GetUserController, mid.JWT([]byte(dbHost)))
	e.GET("users/auth", controllers.GetUserController, mid.JWT([]byte(dbHost)))
	e.POST("/users", controllers.CreateUserController, mid.JWT([]byte(dbHost)))
	e.DELETE("/users/:id", controllers.DeleteUserController, mid.JWT([]byte(dbHost)))
	e.PUT("/users/:id", controllers.UpdateUserController, mid.JWT([]byte(dbHost)))

	return e
}
