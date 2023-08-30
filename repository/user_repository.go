package repository

import (
	"rest-api/model"
)

type UserRepository interface {
	FindAll() []model.User
	FindOneById(id uint64) (model.User, error)
	FindOneByLogin(login string) (model.User, error)
	Create(login, password string) (uint64, error)
	DeleteById(id uint64) error
}
