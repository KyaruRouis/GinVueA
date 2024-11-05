package middleware

import (
	"GinVueA/helper"
	"GinVueA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginAuthCheck 登录认证
func LoginAuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaim, err := helper.AnalyzeToken(c.GetHeader("AccessToken"))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 60403,
				"msg":  "登录过期，请重新登录",
			})
		} else {
			if userClaim.RoleId == 0 {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "非法请求",
				})
			}

			// 判断是否是超管
			sysRole := new(models.SysRole)
			err = models.DB.Model(new(models.SysRole)).Select("is_admin").Where("id = ?", userClaim.RoleId).Find(sysRole).Error
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "网络异常，请重试",
				})
			}

			if sysRole.IsAdmin == 1 {
				userClaim.IsAdmin = true
			} else {
				userClaim.IsAdmin = false
			}

			c.Set("UserClaim", userClaim)
			c.Next()
		}
	}
}
