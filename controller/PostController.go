package controller

import (
	"ginEssential2/common"
	"ginEssential2/model"
	"ginEssential2/response"
	"ginEssential2/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type IPostController interface {
	RestController
	// PageList 前端分页时需要
	PageList(c *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	if !db.Migrator().HasTable(&model.Post{}) {
		err := db.Migrator().CreateTable(&model.Post{})
		if err != nil {
			panic(err)
			return nil
		}
	}
	return PostController{db}
}

func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	if err := c.ShouldBindJSON(&requestPost); err != nil {
		response.Fail(c, nil, "数据验证错误Create")
		return
	}

	//登录用户 user
	user, _ := c.Get("user")

	//创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(c, gin.H{"post": post}, "创建post成功")

}

func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	if err := c.ShouldBindJSON(&requestPost); err != nil {
		response.Fail(c, nil, "数据验证错误Update")
		return
	}

	//获取path中的postId
	postId := c.Params.ByName("id")

	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "文章不存在")
		return
	}

	//当前用户是否为文章作者
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "文章不属于您，请勿操作")
		return
	}

	//更新文章
	if err := p.DB.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(c, gin.H{"post": post}, "更新失败")
		return
	}

	response.Success(c, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(c *gin.Context) {
	//获取path中的postId
	postId := c.Params.ByName("id")

	var post model.Post
	//查询时需要预加载！！关联的其他模型
	if err := p.DB.Preload("Category").Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "文章不存在")
		return
	}

	response.Success(c, gin.H{"post": post}, "查找成功")
}

func (p PostController) Delete(c *gin.Context) {
	//获取path中的postId
	postId := c.Params.ByName("id")

	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "文章不存在")
		return
	}

	//当前用户是否为文章作者
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "文章不属于您，请勿操作")
		return
	}

	if err := p.DB.Delete(&post).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, gin.H{"post": post}, "删除成功")
}

func (p PostController) PageList(c *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "24"))

	//分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	//前端使用分页需知道一共多少条记录，查询记录总条数
	var totalCount int64
	p.DB.Model(model.Post{}).Count(&totalCount)

	response.Success(c, gin.H{"data": posts, "total": totalCount}, "分页成功")

}
