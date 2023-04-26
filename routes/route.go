package route

import (
	"net/http"
	"os"

	"echo_golang/configs"
	"echo_golang/controllers"
	m "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/repositories"
	"echo_golang/services"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	db, _ = configs.InitDB()
	userR = repositories.NewRepository(db)
	userS = services.NewUserService(userR)
	userC = controllers.NewUserController(userS)
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

	c := userC
	e.POST("/login", c.LoginUserController)
	e.GET("/users/:id", c.GetUserController, echojwt.WithConfig(config))
	e.GET("/users", c.GetUsersController, echojwt.WithConfig(config))
	e.GET("users/auth", c.GetUsersController, echojwt.WithConfig(config))
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController, echojwt.WithConfig(config))
	e.PUT("/users/:id", c.UpdateUserController, echojwt.WithConfig(config))

	return e
}
