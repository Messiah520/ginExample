package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

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

	//上传单个文件

	//上传多个文件
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.String(500, "upload err")
		}
		fp, err := file.Open()
		buff := make([]byte, file.Size)
		fp.Read(buff)

		fmt.Println("buff== ", buff)
		//c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, string(buff))
	})

	r.POST("buff", func(c *gin.Context) {

		str := c.PostForm("file")

		fmt.Println([]byte(str))
		c.String(200, "ok")

	})
	//上传多个文件
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("get err %s", err.Error()))
		}

		//获取所有图片
		files := form.File["files"]

		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
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
