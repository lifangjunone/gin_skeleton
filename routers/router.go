package routers

import (
	"github.com/gin-gonic/gin"
)

var engin *gin.Engine

// InitEngine 初始化app
func InitEngine() *gin.Engine {
	engin = gin.New()
	return engin
}

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	// 用户模块API
	InitUserRouter(r)
	// 部门模块API
	InitDeptRouter(r)
}
