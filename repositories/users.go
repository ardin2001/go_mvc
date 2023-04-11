package repositories

import (
	"echo_golang/configs"
	"echo_golang/models"
)

type UserRepositories interface {
	GetUsersRepository() ([]*models.User, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(User models.User) (*models.User, error)
	UpdateRepository(id string, UserBody models.User) (*models.User, error)
	DeleteRepository(id string) error
}

type UserStruct struct {
}

func (us *UserStruct) GetUsersRepository() ([]models.User, error) {
	var users []models.User
	DB, _ := configs.InitDB()
	check := DB.Find(&users).Error

	if check != nil {
		return nil, check
	}

	return users, check
}

func (us *UserStruct) GetUserRepository(id string) (*models.User, error) {
	var user models.User

	DB, _ := configs.InitDB()
	check := DB.First(&user, id).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}

func (us *UserStruct) DeleteRepository(id string) error {
	DB, _ := configs.InitDB()
	check := DB.Delete(&models.User{}, &id).Error

	return check
}

func (us *UserStruct) CreateRepository(user *models.User) (*models.User, error) {
	DB, _ := configs.InitDB()
	check := DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return user, check
}

func (us *UserStruct) UpdateRepository(userId *models.User, id string) (*models.User, error) {
	DB, _ := configs.InitDB()
	var user models.User

	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	if userId.Name != "" {
		user.Name = userId.Name
	}
	if userId.Email != "" {
		user.Email = userId.Email
	}
	if userId.Password != "" {
		user.Password = userId.Password
	}

	check := DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}
