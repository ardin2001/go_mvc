package middleware

import (
	"echo_golang/config"
	"echo_golang/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user models.User
	DB, _ := config.InitDB()
	err := DB.Where("email = ? AND password = ?", username, password).First(&user).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
