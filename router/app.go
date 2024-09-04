package router

import (
	docs "ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine{
	router := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源
	router.Static("/asset", "asset/")
	router.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	r.StaticFS()
	router.LoadHTMLGlob("views/**/*")

	//首页
	router.GET("/", service.GetIndex)
	router.GET("/index",service.GetIndex)

	//用户模块
	router.GET("/user/GetUserList",service.GetUserList)
	router.POST("/user/Register",service.Register)
	router.DELETE("/user/DeleteUser",service.DeleteUser)
	router.PUT("/user/UpdateUser",service.UpdateUser)
	router.POST("/user/Login",service.Login)
	// Websocket发送接受消息测试
	router.GET("/user/SendMessage",service.WebsocketHandler)
	// 发送接受消息
	router.GET("/user/SendUserMessage",service.SendUserMsg)


	return router
}