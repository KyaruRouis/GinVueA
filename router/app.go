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
	// 获取管理员详情信息
	r.GET("/user/detail/:id", service.GetUserDetail)
	// 更新管理员信息
	r.PUT("/user", service.UpdateUser)
	// 删除管理员信息
	r.DELETE("/user/:id", service.DeleteUser)
	// 管理员 end

	// 角色管理 start
	// 角色列表
	r.GET("/role", service.GetRoleList)
	// 添加角色
	r.POST("/role", service.AddRole)
	// 获取角色详情
	r.GET("/role/detail/:id", service.GetRoleDetail)
	// 更新角色信息
	r.PUT("/role", service.UpdateRole)
	// 删除角色信息
	r.DELETE("/role/:id", service.DeleteRole)
	// 修改角色管理员身份
	r.PATCH("/role/:id/:isAdmin", service.PatchRoleAdmin)
	// 角色管理 end

	// 菜单管理 start
	// 菜单列表
	r.GET("/menu", service.GetMenuList)
	// 新增菜单
	r.POST("/menu", service.AddMenu)
	// 更新菜单
	r.PUT("/menu", service.UpdateMenu)
	// 删除菜单
	r.DELETE("/menu/:id", service.DeleteMenu)
	// 菜单管理 end

	return r
}
