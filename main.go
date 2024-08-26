package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig() // 初始化配置文件
	utils.InitMySQL() // 初始化 MySQL 连接
	r := router.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}