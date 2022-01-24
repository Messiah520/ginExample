package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func articleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "article list",
	})
}

func ArticleRouter(e *gin.Engine) {

	e.GET("/articleList", articleList)
}
