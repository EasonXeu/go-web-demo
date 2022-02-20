//+build wireinject

package main

import (
	"github.com/google/wire"
	"go-web-demo/go/controller"
	"go-web-demo/go/dal/mapper"
	"go-web-demo/go/service"
	"gorm.io/gorm"
)

func InitUserController(db *gorm.DB) controller.UserController {
	wire.Build(controller.NewUserController, service.NewUserService, mapper.NewUserMapper)
	return controller.UserController{}
}
