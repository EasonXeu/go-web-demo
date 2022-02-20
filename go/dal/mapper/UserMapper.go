package mapper

import (
	"go-web-demo/go/dal/entity"
	"gorm.io/gorm"
)

type UserMapper struct {
	sqlSession *gorm.DB
}

func (userMapper *UserMapper) CreateUser(user *entity.User) (err error) {
	if err = userMapper.sqlSession.Create(&user).Error; err != nil {
		return err
	}
	return
}

func (userMapper *UserMapper) ListUsers() (userList []*entity.User, err error) {
	if err := userMapper.sqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapper *UserMapper) DeleteUserById(id uint) (err error) {
	var user entity.User
	user.ID = id
	err = userMapper.sqlSession.Delete(&user).Error
	return
}

func (userMapper *UserMapper) GetUserById(id uint) (user *entity.User, err error) {
	if err = userMapper.sqlSession.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapper *UserMapper) UpdateUser(user *entity.User) (err error) {
	err = userMapper.sqlSession.Save(&user).Error
	return
}

func NewUserMapper(sqlSession *gorm.DB) *UserMapper {
	return &UserMapper{sqlSession: sqlSession}
}
