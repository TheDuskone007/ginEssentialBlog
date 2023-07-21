package controller

import (
	"ginEssential2/model"
	"ginEssential2/repository"
	"ginEssential2/response"
	"ginEssential2/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	rep := repository.NewCategoryRepository()
	if !rep.DB.Migrator().HasTable(&model.Category{}) {
		err := rep.DB.Migrator().CreateTable(&model.Category{})
		if err != nil {
			panic(err)
			return nil
		}
	}
	return CategoryController{rep}
}

func (cc CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBindJSON(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称不能为空")
		return
	}

	category, err := cc.Repository.Create(requestCategory.Name)
	if err != nil {
		//response.Fail(c, nil, "创建失败")
		panic(err)
		return
	}

	response.Success(c, gin.H{"category": category}, "创建成功")
}

func (cc CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称不能为空")
		return
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//先查找，再更新分类 map\struct\name value
	targetCategory, err := cc.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	updateCategory, err2 := cc.Repository.Update(*targetCategory, requestCategory.Name)
	if err2 != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")

}

func (cc CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	targetCategory, err := cc.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": targetCategory}, "查找成功")
}

func (cc CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := cc.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
