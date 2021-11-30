package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"helloworld/dao"
	"helloworld/db"
	"helloworld/model"
	"helloworld/response"
	"strconv"
	"time"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db.DB.AutoMigrate(model.Category{})
	return CategoryController{DB: db.DB}
}

func (c CategoryController) AddCategory(ctx *gin.Context) {
	category := &model.Category{
		Name: ctx.PostForm("name"),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	ctx.Bind(&category)
	if category.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	} else {
		category, err := dao.AddCategory(category)
		if nil != err {
			panic(err)
			return
		}
		response.Success(ctx, gin.H{"category":category}, "添加分类成功!")
	}
}

func (c CategoryController) GetCategories(ctx *gin.Context) {
	categories := []*model.Category{}
	categories, err := dao.GetCategories()
	if err != nil {
		response.Fail(ctx, nil, "获取分类失败！")
	}
	response.Success(ctx, gin.H{"categories": categories}, "获取分类成功！")
}

func (c CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	category := &model.Category{
		ID:       uint(id),
		Name:     ctx.PostForm("name"),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	ctx.Bind(&category)
	if category.Name == "" {
		response.Fail(ctx, nil , "数据验证错误，分类名称必填")
	} else {
		dao.UpdateCategory(category)
		response.Success(ctx, gin.H{"category": category}, "修改分类成功")
	}
}

func (c CategoryController) DeleteCategoryByID(ctx *gin.Context) {
	stringID := ctx.PostForm("id")
	id, _ := strconv.Atoi(stringID)
	if err := dao.DeleteCategoryById(uint(id)); nil != err {
		response.Fail(ctx, nil, fmt.Sprint(err))
	}
	response.Success(ctx, nil, "删除分类成功！")
}