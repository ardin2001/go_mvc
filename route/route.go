package route

import (
	"net/http"

	"echo_golang/controllers"
	"echo_golang/middleware"

	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, middleware.TestJWT)
	e.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
