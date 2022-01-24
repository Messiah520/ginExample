package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadRouter(e *gin.Engine) {
	e.POST("upload", upload)
	e.POST("uploads", uploads)
}

//上传多个文件
func uploads(c *gin.Context) {
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
}

func upload(c *gin.Context) {
	str := c.PostForm("file")
	fmt.Println([]byte(str))
	c.String(200, "ok")

}
