package controllers

import (
	"echo_golang/configs"
	middleware "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	LoginUserController(c echo.Context) error
	GetUsersController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
}

type UserStruct struct {
	UserR repositories.UserStruct
}

func (us *UserStruct) GetUserController(c echo.Context) error {
	id := c.Param("id")
	user, check := us.UserR.GetUserRepository(id)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
			"user":    user,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user id " + id,
		"users":   user,
	})
}

func (us *UserStruct) GetUsersController(c echo.Context) error {
	users, check := us.UserR.GetUsersRepository()

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

func (us *UserStruct) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	_, check := us.UserR.CreateRepository(&user)

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

	check := us.UserR.DeleteRepository(id)

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

func (us *UserStruct) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	c.Bind(&user)

	dataUser, check := us.UserR.UpdateRepository(&user, id)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
			"data":    dataUser,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    dataUser,
	})
}

func (us *UserStruct) LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	DB, _ := configs.InitDB()
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
