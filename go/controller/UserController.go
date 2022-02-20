package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/go/common/result"
	"go-web-demo/go/dao/entity"
	"go-web-demo/go/service"
	"gorm.io/gorm"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func (userController *UserController) CreateUser(c *gin.Context) {
	//定义一个User变量
	var user *entity.User
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := userController.userService.CreateUser(user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		result.Error(c, result.InvalidParam, err.Error())
		return
	}
	result.Success(c, user)
	return
}

func (userController *UserController) ListUsers(c *gin.Context) {
	userList, err := userController.userService.ListUsers()
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, userList)
	return
}

func (userController *UserController) GetUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	user, err := userController.userService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		result.Error(c, result.EntityNotExist, err.Error())
		return
	}
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, user)
	return
}
func (userController *UserController) UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	user, err := userController.userService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		result.Error(c, result.EntityNotExist, err.Error())
		return
	}
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	c.BindJSON(&user)
	if err = userController.userService.UpdateUser(user); err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, user)
}

func (userController *UserController) DeleteUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		result.Error(c, result.InvalidParam, "Id is invalid.")
		return
	}
	if err := userController.userService.DeleteUserById(uint(idInInt)); err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.SuccessWithoutData(c)
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}
