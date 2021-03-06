package middleware

import (
	"github.com/fedocx/basicgin/common"
	"github.com/fedocx/basicgin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer"){
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user model.UserInfo
		DB.First(&user,userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg" : "权限不足"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}