package utils

import (
	"echo_golang/config"
	"echo_golang/models"
	"fmt"
)

func UserMigrate() {
	DB, err := config.InitDB()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
