package service

import (
	"go-web-demo/go/dal/entity"
	"go-web-demo/go/dal/mapper"
)

type UserService struct {
	userMapper *mapper.UserMapper
}

func (userService *UserService) CreateUser(user *entity.User) (err error) {
	return userService.userMapper.CreateUser(user)
}

func (userService *UserService) ListUsers() (userList []*entity.User, err error) {
	return userService.userMapper.ListUsers()
}

func (userService *UserService) DeleteUserById(id uint) (err error) {
	return userService.userMapper.DeleteUserById(id)
}

func (userService *UserService) GetUserById(id uint) (user *entity.User, err error) {
	return userService.userMapper.GetUserById(id)
}

func (userService *UserService) UpdateUser(user *entity.User) (err error) {
	return userService.userMapper.UpdateUser(user)
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{userMapper: userMapper}
}
