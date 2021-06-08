package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignForm struct {
	Age        uint8  `json:"age" binding:"gte=18,lte=100"` //18-100之间
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //跨字段比较
}

func main() {
	router := gin.New()
	router.POST("/sign", func(ctx *gin.Context) {
		var signForm SignForm
		if err := ctx.ShouldBind(&signForm); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
}


