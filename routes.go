package main

import (
	"ginEssential2/controller"
	"ginEssential2/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//添加中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//调用中间件保护用户信息接口
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	//都是修改模型，put为替换模型，patch是修改一部分
	//categoryRoutes.PATCH()

	postRoutes := r.Group("/posts")
	//增加获取用户信息的中间件
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.POST("page/list", postController.PageList)

	return r
}
