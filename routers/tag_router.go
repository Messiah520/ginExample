package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func tagListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get mang tags",
	})
}

func TagRouter(e *gin.Engine) {
	e.GET("/tagList", tagListHandler)
}
