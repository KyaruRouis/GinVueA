package main

import (
	"GinVueA/models"
	"GinVueA/router"
)

func main() {

	//初始话gorm.db
	models.NewGormDB()
	r := router.App()
	r.Run(":8015")
}
