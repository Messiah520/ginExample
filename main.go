package main

import (
	"fmt"
	"ginExample/routers"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	//e := casbin.NewEnforcer("./model.conf", "./policy.cvs")
	//auth.Check(e, "dajun", "data1", "read")
	//auth.Check(e, "lizi", "data2", "write")
	//auth.Check(e, "dajun", "data1", "write")
	//auth.Check(e, "dajun", "data2", "read")
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	//api参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+"is"+action)
	})

	//url参数
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintln("hello %s", name))
	})

	//表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")

		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	v1 := r.Group("v1")
	{
		v1.GET("'/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	routers.TagRouter(r)
	routers.ArticleRouter(r)
	routers.UploadRouter(r)

	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
