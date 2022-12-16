package dal

import (
	"go-web-demo/src/dal/dataobject"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserMapperImpl struct {
	SqlSession *gorm.DB
}

func (userMapperImpl *UserMapperImpl) CreateUserDO(user *dataobject.User) (err error) {
	if err = userMapperImpl.SqlSession.Create(&user).Error; err != nil {
		return err
	}
	return
}

func (userMapperImpl *UserMapperImpl) ListUserDOs() (userList []*dataobject.User, err error) {
	if err := userMapperImpl.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapperImpl *UserMapperImpl) DeleteUserDOById(id uint) (err error) {
	var user *dataobject.User
	err = userMapperImpl.SqlSession.Where("name = ?", "test").Delete(&user).Error

	var users []dataobject.User
	userMapperImpl.SqlSession.Clauses(clause.Returning{}).Where("name = ? ", "user1").Delete(&users)
	return
}

func (userMapperImpl *UserMapperImpl) GetUserDOById(id uint) (user *dataobject.User, err error) {
	if err = userMapperImpl.SqlSession.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return
}

func (userMapperImpl *UserMapperImpl) UpdateUserDO(user *dataobject.User) (err error) {
	err = userMapperImpl.SqlSession.Save(&user).Error
	return
}
