package controller

import (
	"fmt"
	"log"
	"net/http"

	"bjzdgt.com/ants/common"
	"bjzdgt.com/ants/dto"
	"bjzdgt.com/ants/model"
	"bjzdgt.com/ants/response"
	"bjzdgt.com/ants/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register for registry user
func Register(ctx *gin.Context) {

	DB := common.GetDB()

	//获取数据
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("tel")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	//判断数据
	if len(name) == 0 {
		name = utils.RandString(9)

	}

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 4221, gin.H{}, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 4222, gin.H{}, "密码需大于6位")
		return
	}

	if !utils.EmailFormatCheck(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 4223, gin.H{}, "email格式错误")
		return
	}
	//处理数据
	var user model.User
	DB.AutoMigrate(&user)

	if UserExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 4224, gin.H{}, "用户已存在")
		return
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 5001, gin.H{}, "密码加密错误")
		return
	}

	DB.Create(&model.User{Name: name, Password: string(hashpassword), Telephone: telephone, Email: email})

	response.Success(ctx, gin.H{}, "ok")

}

// Login is for login
func Login(ctx *gin.Context) {

	DB := common.GetDB()

	//获取数据
	telephone := ctx.PostForm("tel")
	password := ctx.PostForm("password")

	//判断数据
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 4221, gin.H{}, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 4222, gin.H{}, "密码需大于6位")

		return
	}

	//处理数据

	var user model.User
	DB.First(&user, "telephone = ?", telephone)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 4225, gin.H{}, "用户不存在")

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 4226, gin.H{}, "密码错误")
		fmt.Println(err)
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 5001, gin.H{}, "密码解析错误")
		log.Fatal(err)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "ok")

}

//Info return user info
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{
		"user": dto.UserToDto(user.(model.User)),
	}, "ok")
}

//UserExist is a function to determine if the user exist
func UserExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.First(&user, "telephone = ?", tel)
	if user.ID != 0 {
		return true
	}
	return false
}
