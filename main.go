package main

import (
	"fmt"
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
		fmt.Println("1")
		panic(r.Run(":" + port))
	}
	fmt.Println("2")
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
