package models

import "gorm.io/gorm"

// SysLog 日志结构体
type SysLog struct {
	gorm.Model
	Browser     string `gorm:"column:browser;type:varchar(50)" json:"browser"`
	ClassMethod string `gorm:"column:class_method;type:varchar(255);" json:"class_method"`
	HttpMethod  string `gorm:"column:http_method;type:varchar(20)" json:"http_method"`
	Params      string `gorm:"column:params;type:varchar(255);" json:"params"`
	RemoteAddr  string `gorm:"column:remote_addr;type:varchar(255)" json:"remote_addr"`
	RequestUri  string `gorm:"column:request_uri;type:varchar(255)" json:"request_uri"`
	Response    string `gorm:"column:response;type:text(10000);" json:"response"`
	StatusCode  int    `gorm:"column:status_code;type:int(11)" json:"status_code"`
	UseTime     int64  `gorm:"column:use_time;type:varchar(255);" json:"use_time"`
	Country     string `gorm:"column:country;type:varchar(255);" json:"country"`   //国家
	Region      string `gorm:"column:region;type:varchar(255);" json:"region"`     // 区域
	Province    string `gorm:"column:province;type:varchar(255);" json:"province"` // 省份
	City        string `gorm:"column:city;type:varchar(255);" json:"city"`         // 城市
	Isp         string `gorm:"column:isp;type:varchar(255);" json:"isp"`           // 运营商
}

// TableName 设置日志表名称
func (table *SysLog) TableName() string {
	return "sys_log"
}

// GetLogList 获取日志列表
func GetLogList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysLog)).Select("id,browser,class_method,http_method,remote_addr,status_code,use_time,created_at,country,region,province,city,isp")
	if keyword != "" {
		tx.Where("response LIKE ?", "%"+keyword+"%")
	}
	return tx
}
