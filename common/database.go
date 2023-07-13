package common

import (
	"fmt"
	"ginEssential2/model"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := "root"
	password := "123qwe"
	port := "3306"
	dbname := "ginessential"
	charset := "utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, port, dbname, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database, error: " + err.Error())
	}
	//创建数据表
	if !db.Migrator().HasTable(&model.User{}) {
		db.Migrator().CreateTable(&model.User{})
	}
	return db
}

func GetDB() *gorm.DB {
	DB = InitDB()
	return DB
}
