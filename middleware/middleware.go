package middleware

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func MyAuth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method

		sub := "root"

		if ok := e.Enforce(sub, obj, act); ok {
			log.Println("Check successfully")
			c.Next()
		} else {
			log.Println("sorry , Check failed")
			c.Abort()
		}
	}
}
