package middleware

import (
	"ginEssential2/common"
	"ginEssential2/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header（以“Bearer ”开头）
		tokenString := c.GetHeader("Authorization")
		//fmt.Println("tokenString:", tokenString)
		//验证格式 validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "format权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:] //去除“Bearer ”开头
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid { //发生错误或者token无效
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Parse权限不足"})
			c.Abort()
			return
		}

		//通过验证，获取claim中的userId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		//若用户不存在
		if user.ID <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "user权限不足"})
			c.Abort()
			return
		}

		//用户存在，将user信息写入上下文
		c.Set("user", user)

		//Next 应仅在中间件中使用。它执行调用处理程序内链中的挂起处理程序。
		c.Next()
	}
}
