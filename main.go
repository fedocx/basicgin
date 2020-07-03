package main

import (
	"github.com/fedocx/basicgin/common"
	route "github.com/fedocx/basicgin/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main(){
	// init config file,after this program can get config from config file
	InitConfig()

	// init database ,build connection
	common.InitDb()

	//
	server := gin.Default()
	server = route.RouteAdd(server)
	server.Run()
}

func InitConfig(){
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
}


