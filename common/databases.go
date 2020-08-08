package common

import (
	"fmt"
	"ginEsseential/model"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
func init()  {
	if err := InitDB();err !=nil {
		panic(err)
	}
}
func InitDB()  (err error)   {
	db,err =gorm.Open("mysql","root:admin123456@/go_user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Connect Databases Error: %v\n",err)
	}
	db.AutoMigrate(&model.User_info{})
	return
}
func Getdb()  *gorm.DB {
	return db
}