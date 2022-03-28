//+build wireinject

package wire

import (
	"github.com/google/wire"
	"go-web-demo/go/controller"
	"go-web-demo/go/dal/mapper"
	"go-web-demo/go/dal/mapper/mapperimpl"
	"go-web-demo/go/db"
	"go-web-demo/go/service"
	"go-web-demo/go/service/serviceimpl"
)

var MyProviderSet = wire.NewSet(
	db.NewSqlSession,
	//dal level
	wire.Struct(new(mapperimpl.UserMapperImpl), "*"),
	wire.Bind(new(mapper.UserMapper), new(*mapperimpl.UserMapperImpl)),
	//service level
	wire.Struct(new(serviceimpl.UserServiceImpl), "*"),
	wire.Bind(new(service.UserService), new(*serviceimpl.UserServiceImpl)),
	//controller level
	wire.Struct(new(controller.UserController), "*"),
)

func InitUserController() controller.UserController {
	wire.Build(MyProviderSet)
	return controller.UserController{}
}
