package service

import (
	"GinVueA/define"
	"GinVueA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMenuList 获取菜单列表
func GetMenuList(c *gin.Context) {
	Menus(c)
}

// Menus 获取菜单列表数据
func Menus(c *gin.Context) {
	// 登录用户信息
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	data := make([]*MenuReply, 0)
	allMenus := make([]*AllMenu, 0)
	// 根据角色获取所有菜单列表数据
	tx, err := models.GetRoleMenus(userClaim.RoleId, userClaim.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	err = tx.Find(&allMenus).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	data = allMenuToMenuReply(allMenus)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "加载成功",
		"result": data,
	})
}

// 生成树形菜单
func allMenuToMenuReply(allMenu []*AllMenu) []*MenuReply {
	reply := make([]*MenuReply, 0)
	// 一层循环，得到顶层菜单
	for _, v := range allMenu {
		if v.ParentId == 0 {
			reply = append(reply, &MenuReply{
				ID:            v.ID,
				Name:          v.Name,
				WebIcon:       v.WebIcon,
				Sort:          v.Sort,
				Path:          v.Path,
				Level:         v.Level,
				ComponentName: v.ComponentName,
				SubMenus:      getChildrenMenu(v.ID, allMenu),
			})
		}
	}
	return reply
}

// 获取子菜单
func getChildrenMenu(parentId uint, allMenus []*AllMenu) []*MenuReply {
	data := make([]*MenuReply, 0)
	for _, v := range allMenus {
		if v.ParentId == parentId {
			data = append(data, &MenuReply{
				ID:            v.ID,
				Name:          v.Name,
				WebIcon:       v.WebIcon,
				Sort:          v.Sort,
				Path:          v.Path,
				Level:         v.Level,
				ParentId:      v.ParentId,
				ComponentName: v.ComponentName,
				SubMenus:      getChildrenMenu(v.ID, allMenus),
			})
		}
	}
	return data
}

// AddMenu 新增菜单
func AddMenu(c *gin.Context) {
	in := new(AddMenuRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1、保存数据
	err = models.DB.Create(&models.SysMenu{
		ParentId: in.ParentId,
		Name:     in.Name,
		WebIcon:  in.WebIcon,
		Sort:     in.Sort,
		Path:     in.Path,
		Level:    in.Level,
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
		"msg":  "新增成功",
	})

}

// UpdateMenu 更新菜单
func UpdateMenu(c *gin.Context) {
	in := new(UpdateMenuRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	if in.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 更新数据
	err = models.DB.Model(new(models.SysMenu)).Where("id = ?", in.ID).Updates(map[string]interface{}{
		"parent_id":      in.ParentId,
		"name":           in.Name,
		"web_icon":       in.WebIcon,
		"sort":           in.Sort,
		"path":           in.Path,
		"level":          in.Level,
		"component_name": in.ComponentName,
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

// DeleteMenu 删除菜单信息
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}

	// 删除数据库中的数据
	err := models.DB.Where("id = ?", id).Delete(new(models.SysMenu)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})

}
