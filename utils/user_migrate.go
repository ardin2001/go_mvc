package user_migrate

import (
	config_db "echo_golang/config"
	user_models "echo_golang/models"
	"fmt"
)

func UserMigrate() {
	DB, err := config_db.InitDB()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}
	DB.AutoMigrate(&user_models.User{})
}
