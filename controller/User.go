package controller

import (

	//"github.com/fedocx/basicgin/model"
	"github.com/gin-gonic/gin"
)


func UserAdd(c *gin.Context){
	//Username := c.GetPostForm("Username")
	//Password := c.GetPostForm("Password")
	//Telephone := c.GetPostForm("Telephone")
	c.JSON(200,gin.H{
		"message": "用户添加成功",
	})
}