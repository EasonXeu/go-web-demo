//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-web-demo/src/controller"
	"go-web-demo/src/dal"
	"go-web-demo/src/db"
	"go-web-demo/src/service"
)

var MyProviderSet = wire.NewSet(
	db.NewSqlSession,
	//dal level
	wire.Struct(new(dal.UserMapperImpl), "*"),
	wire.Bind(new(dal.UserMapper), new(*dal.UserMapperImpl)),
	//service level
	wire.Struct(new(service.UserServiceImpl), "*"),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	//controller level
	wire.Struct(new(controller.UserController), "*"),
)

func NewUserController() controller.UserController {
	wire.Build(MyProviderSet)
	return controller.UserController{}
}
