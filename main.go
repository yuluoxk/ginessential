package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"log"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRout(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDIr, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDIr + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("read config err:", err)
	}
}
