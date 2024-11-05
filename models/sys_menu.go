package models

import "gorm.io/gorm"

// SysMenu 菜单结构体
type SysMenu struct {
	gorm.Model
	ParentId      uint   `gorm:"column:parent_id;type:int(11);" json:"parent_id"`
	Name          string `gorm:"column:name;type:varchar(100);" json:"name"`
	WebIcon       string `gorm:"column:web_icon;type:varchar(100);" json:"web_icon"`
	Path          string `gorm:"column:path;type:varchar(255);" json:"path"`
	Sort          int    `gorm:"column:sort;type:int(11);default:0;" json:"sort"`                // 排序规则，默认升序，值越小越靠前
	Level         int    `gorm:"column:level;type:tinyint(1);default:0;" json:"level"`           // 菜单等级，{0：目录，1：菜单，2：按钮}
	ComponentName string `gorm:"column:component_name;type:varchar(100);" json:"component_name"` //组件名称
}

// TableName 设置数据表名称
func (table *SysMenu) TableName() string {
	return "sys_menu"
}

// GetMenuList 获取所有菜单数据
func GetMenuList() *gorm.DB {
	tx := DB.Model(new(SysMenu)).Select("id,name,web_icon,sort,path,level,component_name,parent_id").Order("sort ASC")
	return tx
}
