package main

import (
	"ginEssential2/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"os"
)

func main() {
	InitConfig()
	db := common.GetDB()
	if db == nil {
		panic("hi")
	}

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	//go get github.com/spf13/viper
	//go get github.com/dgrijalva/jwt-go
	//go get -u gorm.io/gorm
	//go get -u github.com/go-sql-driver/mysql
	//go get -u github.com/gin-gonic/gin
}

func InitConfig() {

	//获取当前的文本目录
	workDir, _ := os.Getwd()
	//设置要读取的文件名、文件类型、文件路径
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config error")
	}
}
