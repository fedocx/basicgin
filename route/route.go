package route

import (
	"github.com/fedocx/basicgin/middleware"
	"github.com/gin-gonic/gin"
	controller "github.com/fedocx/basicgin/controller"
)


func RouteAdd(server *gin.Engine)*gin.Engine{
	//test
	server.GET("/ping", controller.Ping)
	//user manage
	server.POST("/api/user/register", controller.UserAdd)
	server.POST("/api/user/login", controller.Login)
	server.GET("/api/user/info",middleware.AuthMiddleware(), controller.UserInfo)
	return server
}
