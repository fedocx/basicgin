package controller

import (
	"github.com/fedocx/basicgin/common"
	"github.com/fedocx/basicgin/util"

	//"github.com/fedocx/basicgin/model"
	"github.com/fedocx/basicgin/model"
	"github.com/gin-gonic/gin"
)


func UserAdd(ctx *gin.Context){
	db := common.GetDB()
	requestuser := model.UserInfo{}
	ctx.Bind(&requestuser)

	name := requestuser.Username
	password := requestuser.Password
	telephone := requestuser.Telephone

	if len(telephone) != 11 {
		ctx.JSON(442,gin.H{
			"message": "手机号必须位11位",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(442, gin.H{
			"message": "密码复杂度不符合要求，必须大于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
		requestuser.Username = name
	}
	requestuser.Password = util.Encode(password)
	db.Create(&requestuser)

	ctx.JSON(200,gin.H{
		"message": "用户添加成功",
	})
}

func Login(ctx *gin.Context){
	db := common.GetDB()
	requestuser := model.UserInfo{}
	ctx.Bind(&requestuser)
	var user model.UserInfo
	db.Where("telephone=?", requestuser.Telephone).First(&user)
	if user.ID == 0{
		ctx.JSON(442,gin.H{
			"message":"用户不能为空",
		})
	}
	if user.Password == util.Encode(requestuser.Password) {
		ctx.JSON(200,gin.H{
			"message": "用户登录成功！",
		})
	}
}