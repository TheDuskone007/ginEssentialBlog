package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"type:varchar(110);not null"`
}

func main() {
	db := InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("fail to connect database, error: " + err.Error())
	}
	defer sqlDB.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422 + len(telephone),
				"msg":  "手机号必须为11位",
			})
			return
		}
		fmt.Println("ok")
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码必须大于6位",
			})
			return
		}
		//如果name为空
		if len(name) <= 0 {
			name = RandomString(10)
		}
		fmt.Println(name)

		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已经存在",
			})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
	})
	r.Run(":7080")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjkzxcvbnmqwertyuioASDFGHJKLZXCVBNMQWERTYUIO")
	result := make([]byte, n)
	//初始化随机数生成器
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
	if !db.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
	}
	return db
}
