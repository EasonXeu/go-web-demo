package mapper

import "go-web-demo/go/dal/dataobject"

type UserMapper interface {
	CreateUserDO(user *dataobject.User) (err error)

	ListUserDOs() (userList []*dataobject.User, err error)

	DeleteUserDOById(id uint) (err error)

	GetUserDOById(id uint) (user *dataobject.User, err error)

	UpdateUserDO(user *dataobject.User) (err error)
}
