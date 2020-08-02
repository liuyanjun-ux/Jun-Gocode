package main

import (
	"github.com/garylailai/cateraway-api/handlers/api/database"
	"github.com/garylailai/cateraway-api/handlers/api/router"
)

func main() {
	//初始化Gin服务器
	e := router.InitRouter()
	//初始化gorm
	database.Init()
	e.Run(":1323")
}