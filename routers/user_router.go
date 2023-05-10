package routers

import "github.com/gin-gonic/gin"

func InitUserRouter(r *gin.Engine) {
	userApi := r.Group("/users")
	userApi.GET("/:id")
	userApi.POST("")
	userApi.DELETE("/:id")
	userApi.PATCH("/:id")
}
