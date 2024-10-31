package router

import (
	"GinVueA/middleware"
	"GinVueA/service"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	// 跨域中间件
	r.Use(middleware.Cors())

	// 根据用户名和密码登录路由
	r.POST("/login/password", service.LoginPassword)

	return r
}
