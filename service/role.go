package service

import (
	"GinVueA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetRoleList 获取角色列表数据
func GetRoleList(c *gin.Context) {
	in := &GetRoleListRequest{NewQueryRequest()}
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
		list = make([]*GetRoleListReply, 0)
	)
	err = models.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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
			"total": cnt,
		},
	})

}

// AddRole 新增角色信息
func AddRole(c *gin.Context) {
	in := new(AddRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	//1、判断角色名称是否已经存在
	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常，添加角色失败！",
		})
		return
	}

	// 2、大于0说明角色名称已存在
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "添加角色失败,该角色名称已存在！",
		})
		return
	}

	// 3、保存角色数据
	err = models.DB.Create(&models.SysRole{
		Name:    in.Name,
		Sort:    in.Sort,
		IsAdmin: in.IsAdmin,
		Remarks: in.Remarks,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常，添加角色失败！",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功！",
	})
}

// GetRoleDetail 根据ID获取角色详情信息
func GetRoleDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空！",
		})
		return
	}

	uId, err := strconv.Atoi(id)
	data := new(GetRoleDetailReply)
	sysRole, err := models.GetRoleDetail(uint(uId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常！",
		})
		return
	}

	data.ID = sysRole.ID
	data.Name = sysRole.Name
	data.Sort = sysRole.Sort
	data.IsAdmin = sysRole.IsAdmin
	data.Remarks = sysRole.Remarks
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取数据成功",
		"result": data,
	})

}

// UpdateRole 更新角色信息
func UpdateRole(c *gin.Context) {
	in := new(UpdateRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return

	}

	// 1、判断角色名称是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("id !=? AND name=?", in.ID, in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2、判断是否已经存在
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "更新失败，角色名称已经存在！",
		})
		return
	}
	// 3、更新数据
	err = models.DB.Model(new(models.SysRole)).Where("id = ?", in.ID).Updates(map[string]any{
		"name":     in.Name,
		"is_admin": in.IsAdmin,
		"sort":     in.Sort,
		"remarks":  in.Remarks,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})

}

// DeleteRole 根据ID删除角色信息
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败，ID不能为空！",
		})
		return
	}

	// 删除角色
	err := models.DB.Where("id = ?", id).Delete(new(models.SysRole)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败，数据库异常！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功！",
	})

}

// PatchRoleAdmin 更改管理员身份
func PatchRoleAdmin(c *gin.Context) {
	id := c.Param("id")
	isAdmin := c.Param("isAdmin")
	if id == "" || isAdmin == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参数不能为空",
		})
		return
	}

	// 更改管理员身份
	err := models.DB.Model(new(models.SysRole)).Where("id = ?", id).Update("is_admin", isAdmin).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "更新失败，数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改管理员身份成功",
	})

}
