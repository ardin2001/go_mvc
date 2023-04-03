package route

import (
	"net/http"

	"echo_golang/controllers"
	"echo_golang/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Routers() *echo.Echo {
	e := echo.New()
	middleware.Logger(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/login", controllers.LoginUser)
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController, mid.BasicAuth(middleware.BasicAuthDB))
	e.DELETE("/users/:id", controllers.DeleteUserController, middleware.TestJWT)

	eAuthBasic := e.Group("auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.UpdateUserController)
	eAuthBasic.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
