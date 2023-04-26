package services

import (
	"echo_golang/models"
	"echo_golang/repositories"
)

type UserIntService interface {
	GetUsersService() ([]models.User, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user *models.User) (*models.User, error)
	UpdateService(userId *models.User, id string) (*models.User, error)
	DeleteService(id string) error
}

type userStrService struct {
	userR repositories.UserRepositories
}

func NewUserService(userR repositories.UserRepositories) UserIntService {
	return &userStrService{
		userR: userR,
	}
}

func (u *userStrService) GetUsersService() ([]models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userStrService) GetUserService(id string) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userStrService) CreateService(user *models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *userStrService) UpdateService(userId *models.User, id string) (*models.User, error) {
	user, err := u.userR.UpdateRepository(userId, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userStrService) DeleteService(id string) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
