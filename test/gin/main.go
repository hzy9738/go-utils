package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

var trans ut.Translator

func main() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
	router := gin.New()
	router.Use(MyLogger())
	router.POST("/sign", func(ctx *gin.Context) {
		user := User{}
		_ = ctx.BindJSON(&user)
		fmt.Println(user)
		//var signForm SignForm
		//if err := ctx.ShouldBind(&signForm); err != nil {
		//	errs := err.(validator.ValidationErrors)
		//
		//	ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(trans)})
		//	return
		//}
		ctx.JSON(http.StatusOK, gin.H{
			"name":     user.Name,
			"password": user.Password,
		})
	})

	go func() {
		router.Run(":5000") // listen and serve on 0.0.0.0:8080
	}()

	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("关闭服务...")
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		end := time.Since(t)
		fmt.Printf("耗时：%v\n", end)
		status := c.Writer.Status
		fmt.Println("状态", status)
	}
}

func AuthTokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}
		if token == "" {
			c.Abort()
		}
		c.Next()
	}
}

type SignForm struct {
	Age        uint8  `json:"age" binding:"gte=18,lte=100" label:"年龄"` //18-100之间
	Name       string `json:"name" binding:"required,min=3" label:"姓名"`
	Email      string `json:"email" binding:"required,email" label:"邮箱"`
	Password   string `json:"password" binding:"required" label:"密码"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password" label:"重复密码"` //跨字段比较
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
