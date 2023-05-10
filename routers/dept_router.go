package routers

import "github.com/gin-gonic/gin"

func InitDeptRouter(r *gin.Engine) {
	deptApi := r.Group("/dept")
	deptApi.POST("")
	deptApi.DELETE("/:id")
	deptApi.PATCH("/:id")
	deptApi.GET("/:id")
}
