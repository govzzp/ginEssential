package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context,HttpStatus int, code int,data gin.H,message string)  {
	c.JSON(HttpStatus,gin.H{
		"code":code,
		"data":data,
		"message":message,
	})
}
func Success(c *gin.Context,data gin.H,message string)  {
	Response(c,http.StatusOK,200,data, message)
}
func Fail(c *gin.Context,data gin.H,message string)  {
	Response(c,http.StatusOK,400,data,message)
}