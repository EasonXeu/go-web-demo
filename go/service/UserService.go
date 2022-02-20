package service

import (
	"go-web-demo/go/dal/dataobject"
	"go-web-demo/go/dal/mapper"
)

type UserService interface {
	CreateUser(user *dataobject.User) (err error)

	ListUsers() (userList []*dataobject.User, err error)

	DeleteUserById(id uint) (err error)

	GetUserById(id uint) (user *dataobject.User, err error)

	UpdateUser(user *dataobject.User) (err error)
}

func NewUserService(userMapper mapper.UserMapper) UserService {
	return UserServiceImpl{userMapper: userMapper}
}
