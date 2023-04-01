package route

import (
	"net/http"

	user_controllers "echo_golang/controllers"

	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", user_controllers.GetUserController)
	e.POST("/users", user_controllers.CreateUserController)
	e.DELETE("/users/:id", user_controllers.DeleteUserController)
	e.PUT("/users/:id", user_controllers.UpdateUserController)

	return e
}
