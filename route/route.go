package route

import (

	"github.com/gin-gonic/gin"
	controller "github.com/fedocx/basicgin/controller"
)


func RouteAdd(server *gin.Engine)*gin.Engine{
	//test
	server.GET("/ping", controller.Ping)
	//user manage
	server.GET("/api/user/add", controller.UserAdd)
	return server
}
