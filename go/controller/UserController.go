package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/go/common/result"
	"go-web-demo/go/controller/vo"
	"go-web-demo/go/service"
	"gorm.io/gorm"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

func (userController *UserController) CreateUser(c *gin.Context) {
	//定义一个User变量
	var userVO *vo.UserVO
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&userVO)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建

	user := vo.UserVO2DO(userVO)
	err := userController.UserService.CreateUser(user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		result.Error(c, result.InvalidParam, err.Error())
		return
	}
	result.Success(c, vo.UserDO2VO(user))
	return
}

func (userController *UserController) ListUsers(c *gin.Context) {
	userList, err := userController.UserService.ListUsers()
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, vo.UserDOList2VO(userList))
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
	user, err := userController.UserService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		result.Error(c, result.EntityNotExist, err.Error())
		return
	}
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, vo.UserDO2VO(user))
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
	user, err := userController.UserService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		result.Error(c, result.EntityNotExist, err.Error())
		return
	}
	if err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	c.BindJSON(&user)
	if err = userController.UserService.UpdateUser(user); err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.Success(c, vo.UserDO2VO(user))
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
	if err := userController.UserService.DeleteUserById(uint(idInInt)); err != nil {
		result.Error(c, result.InternalServerError, err.Error())
		return
	}
	result.SuccessWithoutData(c)
}
