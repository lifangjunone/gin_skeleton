package controller

import (
	"gin_skeleton/SkeletonExample/ioc"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/SkeletonExample/service"
	"gin_skeleton/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	UserController *userController
	resp           = common.Resp
	userSvc        = ioc.SvcIocObj.GetSvc("UserService").(service.UserService)
)

type userController struct{}

func newUserController() *userController {
	return &userController{}
}

func (u *userController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := models.NewDefaultUser()
	userSvc.GetUserById(user, id)
	ctx.JSON(http.StatusOK, resp.Success(user))
	return
}

func (u *userController) DeleteUserById() *common.Result {
	return resp.SuccessNotData()
}

func (u *userController) CreateUser(ctx *gin.Context) {
	user := models.NewDefaultUser()
	ctx.ShouldBindJSON(user)
	formUser := user.UserForm
	res := formUser.Validate()
	if res != "" {
		ctx.JSON(http.StatusOK, resp.Error(res))
		return
	}
	userSvc.CreateUser(user)
	ctx.JSON(http.StatusOK, resp.Success(user))
	return
}

func (u *userController) UpdateUserById() *common.Result {
	return resp.SuccessNotData()
}

func init() {
	UserController = newUserController()
}
