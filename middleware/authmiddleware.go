package middleware

import (
	"net/http"
	"strings"

	"bjzdgt.com/ants/common"
	"bjzdgt.com/ants/model"
	"github.com/gin-gonic/gin"
)

// AuthMiddleWare : jwt校验
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足:token 格式错误",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足: token 解析错误",
			})
			return
		}

		var user model.User
		db := common.GetDB()
		db.First(&user, claims.UserID)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 402,
				"msg":  "用户已被删除",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
