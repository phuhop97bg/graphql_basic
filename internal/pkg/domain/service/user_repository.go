package service

import "backend-food/internal/pkg/domain/domain_model/entity"

type UserRepository interface {
	FirstUser(condition entity.Users) (entity.Users, error)
	FindUserList(condition entity.Users) (user []entity.Users, err error)
	FirstUserListWithAnyCondition(query string, condition ...interface{}) (entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	DeleteUser(user entity.Users) error
	UpdateUser(newUser, oldUser entity.Users) (entity.Users, error)
}
