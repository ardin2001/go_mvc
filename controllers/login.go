package controllers

import (
	"echo_golang/config"
	"echo_golang/models"
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	var users models.User
	user := models.User{}
	c.Bind(&user)
	DB, _ := config.InitDB()
	err := DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&users).Error

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "login failed username or password",
		})
	}

	auth := user.Name + ":" + user.Password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "login success",
		"Auth-Basic": enc,
	})
}
