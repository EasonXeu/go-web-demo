package mapper

import (
	"go-web-demo/go/dal/dataobject"
	"gorm.io/gorm"
)

type UserMapperImpl struct {
	sqlSession *gorm.DB
}

func (userMapperImpl UserMapperImpl) CreateUserDO(user *dataobject.User) (err error) {
	if err = userMapperImpl.sqlSession.Create(&user).Error; err != nil {
		return err
	}
	return
}

func (userMapperImpl UserMapperImpl) ListUserDOs() (userList []*dataobject.User, err error) {
	if err := userMapperImpl.sqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapperImpl UserMapperImpl) DeleteUserDOById(id uint) (err error) {
	var user dataobject.User
	user.ID = id
	err = userMapperImpl.sqlSession.Delete(&user).Error
	return
}

func (userMapperImpl UserMapperImpl) GetUserDOById(id uint) (user *dataobject.User, err error) {
	if err = userMapperImpl.sqlSession.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapperImpl UserMapperImpl) UpdateUserDO(user *dataobject.User) (err error) {
	err = userMapperImpl.sqlSession.Save(&user).Error
	return
}
