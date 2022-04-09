package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
	}

	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "分类不存在")
	}

	// 更新分类
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if err := c.DB.First(&category, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "分类不存在")
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
	}
	response.Success(ctx, nil, "")
}
