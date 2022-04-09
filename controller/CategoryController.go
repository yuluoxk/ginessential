package controller

import (
	"ginEssential/model"
	"ginEssential/repository"
	"ginEssential/response"
	"ginEssential/vo"
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
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
		return
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "")
}
