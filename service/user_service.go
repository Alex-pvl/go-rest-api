package service

import (
	"rest-api/model"
	"rest-api/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) GetAllUsers() []model.User {
	return us.UserRepository.FindAll()
}

func (us *UserService) GetUserById(id uint64) (model.User, error) {
	return us.UserRepository.FindOneById(id)
}

func (us *UserService) GetUserByLogin(login string) (model.User, error) {
	return us.UserRepository.FindOneByLogin(login)
}

func (us *UserService) AddUser(login, password string) (uint64, error) {
	return us.UserRepository.Create(login, password)
}

func (us *UserService) DeleteUser(id uint64) error {
	return us.UserRepository.DeleteById(id)
}
