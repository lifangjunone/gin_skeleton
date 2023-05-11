package controller

import (
	"gin_skeleton/SkeletonExample/ioc"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/SkeletonExample/service"
	"gin_skeleton/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	intId, _ := strconv.Atoi(id)
	user := models.NewDefault()
	userSvc.GetUserById(user, intId)
	ctx.JSON(http.StatusOK, resp.Success(user))
	return
}

func (u *userController) DeleteUserById() *common.Result {
	return resp.SuccessNotData()
}

func (u *userController) CreateUser(ctx *gin.Context) {
	user := models.NewDefault()
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
