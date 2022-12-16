package service

import (
	"go-web-demo/src/dal/dataobject"
)

type UserService interface {
	CreateUser(user *dataobject.User) (err error)

	ListUsers() (userList []*dataobject.User, err error)

	DeleteUserById(id uint) (err error)

	GetUserById(id uint) (user *dataobject.User, err error)

	UpdateUser(user *dataobject.User) (err error)
}
