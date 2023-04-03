package controllers

import (
	"echo_golang/config"
	"echo_golang/middleware"
	"echo_golang/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []models.User

	DB, _ := config.InitDB()
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
	user := models.User{}
	c.Bind(&user)
	DB, _ := config.InitDB()
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

func DeleteUserController(c echo.Context) error {
	fmt.Println("punten delete data")
	id := c.Param("id")
	DB, _ := config.InitDB()

	check := DB.Delete(&models.User{}, &id).Error
	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    "data id " + id + " berhasil dihapus",
	})

}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	DB, _ := config.InitDB()
	user := models.User{}

	DB.First(&user, id)
	// new_id, _ := strconv.Atoi(id)
	// user.ID = new_id
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "success",
			"data":    "ERROR INPUT",
		})
	}
	DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})

}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	DB, _ := config.InitDB()
	err := DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "login failed username or password",
			"error":   err.Error(),
		})
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	userresponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"users":   userresponse,
	})
}
