package service

import (
	"go-web-demo/src/dal"
	"go-web-demo/src/dal/dataobject"
)

type UserServiceImpl struct {
	UserMapper dal.UserMapper
}

func (userServiceImpl UserServiceImpl) CreateUser(user *dataobject.User) (err error) {
	return userServiceImpl.UserMapper.CreateUserDO(user)
}

func (userServiceImpl UserServiceImpl) ListUsers() (userList []*dataobject.User, err error) {
	return userServiceImpl.UserMapper.ListUserDOs()
}

func (userServiceImpl *UserServiceImpl) DeleteUserById(id uint) (err error) {
	return userServiceImpl.UserMapper.DeleteUserDOById(id)
}

func (userServiceImpl *UserServiceImpl) GetUserById(id uint) (user *dataobject.User, err error) {
	return userServiceImpl.UserMapper.GetUserDOById(id)
}

func (userServiceImpl *UserServiceImpl) UpdateUser(user *dataobject.User) (err error) {
	return userServiceImpl.UserMapper.UpdateUserDO(user)
}
