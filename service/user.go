package service

import (
	"GinVueA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList 获取管理员列表数据
func GetUserList(c *gin.Context) {
	in := &GetUserListRequest{NewQueryRequest()}
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
		list = make([]*GetUserListReply, 0)
	)
	err = models.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载数据成功",
		"result": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

// AddUser 新增管理员信息
func AddUser(c *gin.Context) {
	in := new(AddUserRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	// 1.判断用户名已经存在
	var cnt int64
	err = models.DB.Model(new(models.SysUser)).Where("username = ?", in.Username).Count(&cnt).Error

	// 大于0说明用户已经存在
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "添加失败，用户名已经存在",
		})
		return
	}
	// 2.保存数据
	err = models.DB.Create(&models.SysUser{
		UserName: in.Username,
		PassWord: in.Password,
		Phone:    in.Phone,
		Email:    in.Email,
		Remarks:  in.Remarks,
		RoleId:   in.RoleId,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "添加失败，数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
	})
}

// GetUserDetail 根据ID获取管理员详情信息
func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		return
	}

	uId, err := strconv.Atoi(id)
	data := new(GetUserDetailReply)
	// 1.获取管理员信息
	sysUser, err := models.GetUserDetail(uint(uId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	data.ID = sysUser.ID
	data.Remarks = sysUser.Remarks
	data.Phone = sysUser.Phone
	data.Username = sysUser.UserName
	data.Email = sysUser.Email
	data.Password = sysUser.PassWord
	data.RoleId = sysUser.RoleId

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取数据成功",
		"result": data,
	})

}

// UpdateUser 修改管理员信息
func UpdateUser(c *gin.Context) {
	in := new(UpdateUserRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1、判断用户名是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysUser)).Where("id !=? AND username = ?", in.ID, in.Username).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}

	// 2、修改数据
	err = models.DB.Model(new(models.SysUser)).Where("id = ?", in.ID).Updates(map[string]any{
		"password": in.Password,
		"username": in.Username,
		"phone":    in.Phone,
		"email":    in.Email,
		"remarks":  in.Remarks,
		"role_id":  in.RoleId,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})

}

// DeleteUser 删除管理员信息
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败，id不能为空",
		})
		return
	}

	// 删除管理员
	err := models.DB.Where("id = ?", id).Delete(new(models.SysUser)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败，数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// AllRoleListReply 返回角色数据的结构体
type AllRoleListReply struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
