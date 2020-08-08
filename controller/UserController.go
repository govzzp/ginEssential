package controller

import (
	"ginEsseential/common"
	"ginEsseential/dto"
	"ginEsseential/model"
	"ginEsseential/response"
	"ginEsseential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(c *gin.Context)  {
	db := common.Getdb()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"Telephone number must be 11 numbers")
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"Telephone number must be 11 numbers",
		//})
		return
	}
	if len(password) <=6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"Password must more than 6 words")
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"Password must more than 6 words",
		//})
		return
	}

	//判断手机号是否存在
	var user model.User_info
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"Unknown User")
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"Unknown User",
		//})
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil{
		response.Response(c,http.StatusBadRequest,400,nil,"密码错误")
		//c.JSON(http.StatusBadRequest,gin.H{
		//	"code":400,
		//	"message":"密码错误",
		//})
		return
	}


	//发放token
	token ,err :=common.ReleaseToken(user)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"server error")
		//c.JSON(http.StatusInternalServerError,gin.H{
		//	"code":500,
		//	"msg":"server error",
		//	//log.Printf("toker gernerate error : %v",err)
		//})
		return
	}
	//返回结果
	//c.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	//	"message":"登录成功",
	//}
	response.Success(c,gin.H{"token":token},"登录成功")
}
func Register(c *gin.Context) {
	db := common.Getdb()
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//验证数据
	if len(telephone) != 11 {
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"Telephone number must be 11 numbers",
		//})
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"Telephone number must be 11 numbers")
		return
	}
	if len(password) <=6 {
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"Password must more than 6 words",
		//})
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"Password must more than 6 words")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isTelephoneExist(db,telephone) {
		//c.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"message":"User is Already exist Please login",
		//})
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"User is Already exist Please login")
		return
	}

	hasedPassword ,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		//c.JSON(http.StatusInternalServerError,gin.H{
		//	"code":500,
		//	"message":"Internal Server Error",
		//})
		response.Response(c,http.StatusInternalServerError,500,nil,"Internal Server Error")
	}
	newUser:=model.User_info{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	db.Create(&newUser)
	log.Println(name,password,telephone)
	//c.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"massage":"Register Successful",
	//})
	response.Success(c,nil,"Register Successful")
}
func isTelephoneExist(db *gorm.DB,telephone string) bool {
	var user model.User_info
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func Info(c *gin.Context)  {
	user , _ := c.Get("user")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{"user":dto.ToUserDto(user.(model.User_info))},
	})
}