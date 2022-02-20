package service

import (
	"go-web-demo/go/dal/dataobject"
	"go-web-demo/go/dal/mapper"
)

type UserServiceImpl struct {
	userMapper mapper.UserMapper
}

func (userServiceImpl UserServiceImpl) CreateUser(user *dataobject.User) (err error) {
	return userServiceImpl.userMapper.CreateUserDO(user)
}

func (userServiceImpl UserServiceImpl) ListUsers() (userList []*dataobject.User, err error) {
	return userServiceImpl.userMapper.ListUserDOs()
}

func (userServiceImpl UserServiceImpl) DeleteUserById(id uint) (err error) {
	return userServiceImpl.userMapper.DeleteUserDOById(id)
}

func (userServiceImpl UserServiceImpl) GetUserById(id uint) (user *dataobject.User, err error) {
	return userServiceImpl.userMapper.GetUserDOById(id)
}

func (userServiceImpl UserServiceImpl) UpdateUser(user *dataobject.User) (err error) {
	return userServiceImpl.userMapper.UpdateUserDO(user)
}
