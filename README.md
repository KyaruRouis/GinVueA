# 环境

GO 1.23.2

mysql 5.7

models目录中的init.go文件配置数据库。

# 项目搭建

## GIN框架

go get -u github.com/gin-gonic/gin

## GORM框架

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

## jwt-go

go get -u github.com/dgrijalva/jwt-go

## 发送邮件插件

go get github.com/jordan-wright/email

## sessions

go get github.com/gin-contrib/sessions

## ip2region插件

go get github.com/lionsoul2014/ip2region/binding/golang

# go run启动

go run main.go
