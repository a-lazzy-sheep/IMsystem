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
	router.GET("/toRegister", service.ToRegister)
	router.GET("/toChat", service.ToChat)
	router.GET("/chat", service.Chat)
	router.POST("/searchFriends", service.SearchFriends)

	//用户模块
	router.GET("/user/GetUserList",service.GetUserList)
	router.POST("/user/Register",service.Register)
	router.DELETE("/user/DeleteUser",service.DeleteUser)
	router.PUT("/user/UpdateUser",service.UpdateUser)
	router.POST("/user/Login",service.Login)
	router.POST("/user/find", service.FindByID)

	// Websocket发送接受消息测试
	router.GET("/user/SendMessage",service.WebsocketHandler)
	//发送消息
	router.GET("/user/sendMsg", service.SendMsg)
	// 发送接受消息
	router.GET("/user/SendUserMessage",service.SendUserMsg)
	//添加好友
	router.POST("/contact/addfriend", service.AddFriend)
	//加载群列表
	router.POST("/contact/loadcommunity", service.LoadCommunity)
	//创建群
	router.POST("/contact/createCommunity", service.CreateCommunity)
	//加入群
	router.POST("/contact/joinCommunity", service.JoinCommunity)
	//退出群
	//上传文件
	router.POST("/attach/upload", service.Upload)

	// 缓存消息
	router.POST("/user/redisMsg", service.RedisMsg)


	return router
}