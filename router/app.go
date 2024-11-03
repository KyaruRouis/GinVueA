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
	// 管理员 start
	// 管理列表
	r.GET("/user", service.GetUserList)
	// 新增管理员信息
	r.POST("/user", service.AddUser)

	// 管理员 end

	return r
}
