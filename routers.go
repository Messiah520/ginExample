package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}

//Json数据绑定和解绑
type LoginJson struct {
	User     string
	Password string
}

func loginJson(c *gin.Context) {
	var json LoginJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

//表单数据解析和绑定
type LoginForm struct {
	User     string
	Password string
}

func loginForm(c *gin.Context) {
	var form LoginForm

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

//URI数据解析和绑定
type LoginUrl struct {
	User    string
	Pssword string
}

func loginUrl(c *gin.Context) {
	var login LoginUrl
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

//结构体验证
type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func personDemo(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}

//自定义中间件
type Person2 struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 2、在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

//自定义校验方法
func nameNotNullAndAdmin(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if value, ok := field.Interface().(string); ok {
		return value != "" && !("5lmh" == value)
	}
	return true
}
