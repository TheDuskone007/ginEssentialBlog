package main

import (
	"ginEssential2/common"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

func main() {
	db := common.GetDB()
	if db == nil {
		panic("hi")
	}

	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":7080")
}
