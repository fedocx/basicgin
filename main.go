package main

import (
	route "github.com/fedocx/basicgin/route"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	server = route.RouteAdd(server)
	server.Run()
}


