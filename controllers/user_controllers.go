package controllers

import (
	"echo_golang/config"
	"echo_golang/middleware"
	"echo_golang/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	LoginUserController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
}

type UserStruct struct {
}

func (us *UserStruct) GetUserController(c echo.Context) error {
	var users []models.User
	claim, _ := middleware.GetClaims(c)
	DB, _ := config.InitDB()
	fmt.Println(DB)
	check := DB.Find(&users).Error
	fmt.Println(check)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
		"auth-db": claim,
	})

}

func (us *UserStruct) CreateUserController(c echo.Context) error {
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

func (us *UserStruct) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	DB, _ := config.InitDB()

	data, _ := middleware.Restricted(c)
	check := DB.Delete(&models.User{}, &id).Error
	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success",
		"data":      "data id " + id + " berhasil dihapus",
		"data-auth": data,
	})

}

func (us *UserStruct) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	DB, _ := config.InitDB()
	user := models.User{}

	DB.First(&user, id)
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

func (us *UserStruct) LoginUserController(c echo.Context) error {
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
