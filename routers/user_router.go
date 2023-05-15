package routers

import "github.com/gin-gonic/gin"
import "gin_skeleton/SkeletonExample/controller"

var (
	userController = controller.UserController
)

func InitUserRouter(r *gin.Engine) {
	userApi := r.Group("/users")
	userApi.GET("/:id", userController.GetUserById)
	userApi.POST("", userController.CreateUser)
	userApi.POST("/query", userController.GetByField)
	userApi.DELETE("/:id")
	userApi.PATCH("/:id")
}
