package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(50);" json:"userName"`
	PassWord  string `gorm:"column:password;type:varchar(36);" json:"passWord"`
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);" json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);" json:"wxOpenId"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
}

// TableName 设置表名称
func (table *SysUser) TableName() string {
	return "sys_user"
}

// GetUserByUsernamePassword 根据用户名和密码查询数据
func GetUserByUsernamePassword(username string, password string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("username = ? AND password = ?", username, password).First(data).Error
	return data, err
}
