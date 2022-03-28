package serviceimpl

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-web-demo/go/dal/dataobject"
	"go-web-demo/go/service/serviceimpl/mocks"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userMapper := &mocks.UserMapper{}

	userMapper.EXPECT().CreateUserDO(mock.Anything).Return(nil)

	userServiceImpl := UserServiceImpl{
		UserMapper: userMapper,
	}

	user := dataobject.User{ID: 123}
	err := userServiceImpl.CreateUser(&user)
	assert.Nil(t, err)

}
