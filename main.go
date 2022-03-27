package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRout(r)
	panic(r.Run())
}
