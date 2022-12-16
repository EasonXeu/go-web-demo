package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-demo/src/controller"
)

var userController *controller.UserController

func SetRouter() *gin.Engine {
	r := gin.Default()

	userController := NewUserController()
	/**
	用户User路由组
	*/
	userGroup := r.Group("/v1")
	userGroup.Use(MyMiddleware())

	{
		//增加用户User
		userGroup.POST("/users", userController.CreateUser)
		//查看所有的User
		userGroup.GET("/users", userController.ListUsers)

		userGroup.GET("/users/:id", userController.GetUserById)
		//修改某个User
		userGroup.PUT("/users/:id", userController.UpdateUser)
		//删除某个User
		userGroup.DELETE("/users/:id", userController.DeleteUserById)
	}

	return r
}

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(fmt.Sprintf("log-%s-%s", c.Request.Method, c.FullPath()))
	}
}
