package controller

import (
	"gin_skeleton/SkeletonExample/service"
	"gin_skeleton/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	UserController *userController
	resp           = common.Resp
	userService    = service.UserSvc
)

type userController struct{}

func newUserController() *userController {
	return &userController{}
}

func (u *userController) GetUserById(context *gin.Context) {
	id := context.Param("id")
	intId, _ := strconv.Atoi(id)
	user := userService.GetUserById(intId)
	context.JSON(http.StatusOK, resp.Success(user))
	return
}

func (u *userController) DeleteUserById() *common.Result {
	return resp.SuccessNotData()
}

func (u *userController) CreateUserById() *common.Result {
	return resp.SuccessNotData()
}

func (u *userController) UpdateUserById() *common.Result {
	return resp.SuccessNotData()
}

func init() {
	UserController = newUserController()
}
