package main

import (
	"ginEsseential/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	r := gin.Default()
	r = routes.ConnectRouter(r)
	panic(r.Run(":8080"))
}
