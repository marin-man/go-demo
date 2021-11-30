package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"helloworld/dao"
	"helloworld/db"
	"helloworld/model"
	"helloworld/response"
	"strconv"
)

type IPostController interface {
	RestPostController
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db.DB.AutoMigrate(model.Post{})
	return PostController{DB: db.DB}
}

func (p PostController) AddPost(ctx *gin.Context) {
	userID := ctx.PostForm("user_id")
	categoryID := ctx.PostForm("category_id")
	title := ctx.PostForm("title")
	headImg := ctx.PostForm("head_img")
	content := ctx.PostForm("content")
	uuserID, _ := strconv.Atoi(userID)
	ucategoryID, _ := strconv.Atoi(categoryID)
	post := &model.Post{
		UserId:     uint(uuserID),
		CategoryId: uint(ucategoryID),
		Title:      title,
		HeadImg:    headImg,
		Content:    content,
	}
	dao.AddPost(post)
}

func (p PostController) GetPost(ctx *gin.Context) {
	posts := []*model.Post{}
	posts, err := dao.GetPost()
	if err != nil {
		response.Fail(ctx, nil, "获取文章失败！")
	}
	response.Success(ctx, gin.H{"posts": posts}, "获取文章成功！")
}

func (p PostController) UpdatePost(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")

	var post model.Post
	if db.DB.Where("id=?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx,  nil, "post not exist!")
		return
	}
	user, _ := ctx.Get("user") // 判断当前用户是否为文章的作者
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章不属于您，请不要非法操作！")
		return
	}
	err := db.DB.Model(&post).Update(post).Error // Update post
	if err != nil {
		response.Fail(ctx, nil, "Update fail")
	}
	response.Success(ctx, gin.H{"post": post}, "Update success!")
}

func (p PostController) DeletePostByID(ctx *gin.Context) {
	postId := ctx.Params.ByName("id") // Get the path id
	var post model.Post
	if db.DB.Where("id=?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "Post not exist!")
		return
	}
	user, _ := ctx.Get("user") // 获取登录用户
	userId := user.(model.User).ID
	if userId != post.UserId { // 判断当前用户是否为文章的作者
		response.Fail(ctx, nil, "文章不属于您，请勿非法操作！")
		return
	}
	db.DB.Delete(&post)
	response.Success(ctx, gin.H{"post": post}, "Delete success!")
}