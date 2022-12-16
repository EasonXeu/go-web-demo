package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/src/controller/vo"
	"go-web-demo/src/service"
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
	if err := c.ShouldBindJSON(&userVO); err != nil {
		Abort(c, InvalidParam, err.Error())
		return
	}

	user := vo.UserVO2DO(userVO)

	if err := userController.UserService.CreateUser(user); err != nil {
		Abort(c, InvalidParam, err.Error())
		return
	}
	Success(c, vo.UserDO2VO(user))
	return
}

func (userController *UserController) ListUsers(c *gin.Context) {
	userList, err := userController.UserService.ListUsers()
	if err != nil {
		Abort(c, InternalServerError, err.Error())
		return
	}
	Success(c, vo.UserDOList2VO(userList))
	return
}

func (userController *UserController) GetUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	user, err := userController.UserService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		Abort(c, EntityNotExist, err.Error())
		return
	}
	if err != nil {
		Abort(c, InternalServerError, err.Error())
		return
	}
	Success(c, vo.UserDO2VO(user))
	return
}
func (userController *UserController) UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	user, err := userController.UserService.GetUserById(uint(idInInt))
	if err == gorm.ErrRecordNotFound {
		Abort(c, EntityNotExist, err.Error())
		return
	}
	if err != nil {
		Abort(c, InternalServerError, err.Error())
		return
	}
	c.BindJSON(&user)
	if err = userController.UserService.UpdateUser(user); err != nil {
		Abort(c, InternalServerError, err.Error())
		return
	}
	Success(c, vo.UserDO2VO(user))
}

func (userController *UserController) DeleteUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		Abort(c, InvalidParam, "Id is invalid.")
		return
	}
	if err := userController.UserService.DeleteUserById(uint(idInInt)); err != nil {
		Abort(c, InternalServerError, err.Error())
		return
	}
	SuccessWithoutData(c)
}
