package user_controllers

import (
	config_db "echo_golang/config"
	user_models "echo_golang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []user_models.User

	DB, _ := config_db.InitDB()
	check := DB.Find(&users).Error

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})

}

func CreateUserController(c echo.Context) error {
	user := user_models.User{}
	c.Bind(&user)

	DB, _ := config_db.InitDB()
	check := DB.Save(&user).Error

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})

}
