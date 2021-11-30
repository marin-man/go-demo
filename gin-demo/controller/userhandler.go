package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"helloworld/common"
	"helloworld/dao"
	"helloworld/db"
	"helloworld/dto"
	"helloworld/model"
	"helloworld/response"
	"helloworld/util"
	"log"
	"net/http"
)

// Register 注册
func Register(ctx *gin.Context) {
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.Param("password")
	phone := ctx.PostForm("phone")

	// 数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "The phone num must be 11 digits",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":422,
			"msg": "Password cannot be less than 6 digits!",
		})
		return
	}
	if len(username) == 0 {
		username = util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User exist!")
		return
	}

	BasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if nil != err {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Based password error!")
		return
	}
	user := model.User{
		Username: username,
		Password: string(BasedPassword),
		Phone: phone,
	}
	db.DB.Create(&user)
	response.Success(ctx, nil, "Register success!")
}

func Login(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	if phone == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Phone num not null!")
		return
	}
	if password == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password num not null!")
		return
	}
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Phone num must 11 digits!")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password len not less 6 bigits!")
		return
	}

	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User not exist!")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); nil != err {
		response.Fail(ctx, nil, "password err!")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "System err!")
		log.Printf("token generate error: %v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}