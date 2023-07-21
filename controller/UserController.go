package controller

import (
	"fmt"
	"ginEssential2/common"
	"ginEssential2/dto"
	"ginEssential2/model"
	"ginEssential2/response"
	"ginEssential2/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	////使用map获取请求得参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(c.Request.Body).Decode(&requestMap)

	////获取参数
	//name := c.PostForm("name")
	//telephone := c.PostForm("telephone")
	//password := c.PostForm("password")

	//使用ShouldBind绑定结构体获取参数
	var requestUser = model.User{}
	err := c.ShouldBindJSON(&requestUser)
	if err != nil {
		response.Response(c, http.StatusBadGateway, 502, nil, err.Error())
	}
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//如果name为空
	if len(name) <= 0 {
		name = util.RandomString(10)
	}
	fmt.Println(name)

	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	//创建用户&&密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "注册并登陆成功")
	//response.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	//使用ShouldBind绑定结构体获取参数
	var requestUser = model.User{}
	err := c.ShouldBindJSON(&requestUser)
	if err != nil {
		response.Response(c, http.StatusBadGateway, 502, nil, err.Error())
	}
	//name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//telephone := c.PostForm("telephone")
	//password := c.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//判断手机号是否存在
	db := common.GetDB()
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID <= 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")

}

func Info(c *gin.Context) {
	//获取用户信息
	user, _ := c.Get("user")

	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取用户信息成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
