// gin 初始化   gin.default   返回结构体指针
package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode("release")  // 设置模式 release|debug
	r := gin.Default()   // 默认启动方式，包含 Logger、Recovery 中间件   r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code" :200,
			"message": "pong",
		})
	})

	r.POST("/post", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code":200,
			"message":"test-post",
		})
	})

	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		//v2.POST("/submit", submitEndpoint)
		//v2.POST("/read", readEndpoint)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
func loginEndpoint(content *gin.Context)  {

	content.JSON(200,gin.H{
		"code":200,
		"message":"loginEndpoint",
	})
}
