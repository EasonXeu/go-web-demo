package serviceimpl

import (
	"go-web-demo/go/dal/dataobject"
	"go-web-demo/go/dal/mapper"
)

type UserServiceImpl struct {
	UserMapper mapper.UserMapper
}

func (userServiceImpl *UserServiceImpl) CreateUser(user *dataobject.User) (err error) {
	return userServiceImpl.UserMapper.CreateUserDO(user)
}

func (userServiceImpl *UserServiceImpl) ListUsers() (userList []*dataobject.User, err error) {
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
