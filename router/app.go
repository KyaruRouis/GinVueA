package router

import (
	"GinVueA/define"
	"GinVueA/middleware"
	"GinVueA/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	// 添加跨域中间件
	r.Use(middleware.Cors())

	// 日志中间件
	r.Use(middleware.LoggerToDb())

	// 使用session中间件
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Secure:   true,
		SameSite: 4,
		Path:     "/",
		MaxAge:   3600,
	})
	r.Use(sessions.Sessions("admin-session", store))

	// 设置静态资源路由
	r.Static("/uploadFile", define.StaticResource)

	// 根据用户名和密码登录路由
	r.POST("/login/password", service.LoginPassword)

	// 登录信息认证
	loginAuth := r.Group("/").Use(middleware.LoginAuthCheck())

	// 上传图片文件
	loginAuth.POST("/upload/file", service.UploadFile)

	// 管理员 start
	// 管理列表
	loginAuth.GET("/user", service.GetUserList)
	// 新增管理员信息
	loginAuth.POST("/user", service.AddUser)
	// 获取管理员详情信息
	loginAuth.GET("/user/detail/:id", service.GetUserDetail)
	// 更新管理员信息
	loginAuth.PUT("/user", service.UpdateUser)
	// 删除管理员信息
	loginAuth.DELETE("/user/:id", service.DeleteUser)

	// 更新个人信息
	loginAuth.PUT("/user/updateInfo", service.UpdateInfo)

	// 发送邮件
	loginAuth.GET("/user/sendEmail", service.SendEmail)

	// 校验邮箱验证码
	loginAuth.GET("/user/verifyCode", service.VerifyCode)

	// 更新绑定邮箱
	loginAuth.PUT("/user/updateEmail", service.UpdateEmail)

	// 更改个人密码
	loginAuth.PUT("/user/updatePwd", service.UpdatePwd)

	// 管理员 end

	// 角色管理 start
	loginAuth.GET("/role", service.GetRoleList)
	// 添加角色
	loginAuth.POST("/role", service.AddRole)
	// 获取角色详情
	loginAuth.GET("/role/detail/:id", service.GetRoleDetail)
	// 更新角色信息
	loginAuth.PUT("/role", service.UpdateRole)
	// 删除角色信息
	loginAuth.DELETE("/role/:id", service.DeleteRole)
	// 修改角色管理员身份
	loginAuth.PATCH("/role/:id/:isAdmin", service.PatchRoleAdmin)
	// 获取所有角色
	loginAuth.GET("/role/all", service.AllRole)
	// 角色管理 end

	// 菜单管理 start
	// 菜单列表
	loginAuth.GET("/menu", service.GetMenuList)
	// 新增菜单
	loginAuth.POST("/menu", service.AddMenu)
	// 更新菜单
	loginAuth.PUT("/menu", service.UpdateMenu)
	// 删除菜单
	loginAuth.DELETE("/menu/:id", service.DeleteMenu)
	// 菜单管理 end

	// 日志管理 start
	loginAuth.GET("/log", service.GetLogList)
	// 日志管理 end

	return r
}
