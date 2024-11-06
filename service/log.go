package service

import (
	"GinVueA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLogList 获取日志列表
func GetLogList(c *gin.Context) {
	in := &GetLogListRequest{NewQueryRequest()}
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		cnt  int64
		list = make([]*GetLogListReply, 0)
	)

	err = models.GetLogList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"result": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}
