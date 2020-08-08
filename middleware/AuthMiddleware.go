package middleware

import (
	"ginEsseential/common"
	"ginEsseential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		//获取Authorization Header
		tokenString := c.GetHeader("Authorization")
		//验证格式
		if tokenString == ""||!strings.HasPrefix(tokenString,"Bearer"){
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"message":"权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token ,claims,err := common.ParseToken(tokenString)
		if err != nil||!token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"message":"权限不足",
			})
			c.Abort()
			return
		}
		//验证通过后，获取Claims中的UserID
		userID := claims.UserID
		db := common.Getdb()
		var user model.User_info
		db.First(&user,userID)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"message":"权限不足",
			})
			c.Abort()
			return
		}
		//用户存在将User信息写入上下文
		c.Set("user",user)
		c.Next()
	}
}