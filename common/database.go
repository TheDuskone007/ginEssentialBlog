package common

import (
	"fmt"
	"ginEssential2/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasoure.user")
	password := viper.GetString("datasoure.password")
	host := viper.GetString("datasoure.host")
	port := viper.GetString("datasoure.port")
	database := viper.GetString("datasoure.database")
	charset := viper.GetString("datasoure.charset")
	loc := viper.GetString("datasoure.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		user, password, host, port, database, charset, url.QueryEscape(loc))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database, error: " + err.Error())
	}
	//迁移user数据表，判断是否有，若无则创建
	if !db.Migrator().HasTable(&model.User{}) {
		db.Migrator().CreateTable(&model.User{})
	}
	return db
}

func GetDB() *gorm.DB {
	DB = InitDB()
	return DB
}
