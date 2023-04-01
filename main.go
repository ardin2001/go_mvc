package main

import (
	"net/http"

	user_controllers "echo_golang/controllers"
	user_migrate "echo_golang/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	user_migrate.UserMigrate()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", user_controllers.GetUserController)
	e.POST("/users", user_controllers.CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
