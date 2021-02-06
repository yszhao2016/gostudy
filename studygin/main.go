package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"go-gin-api/app/config"
)

func main() {
	fmt.Println("GIN Framework  Study")

	//gin.Default()
	gin.SetMode("debug")
	engine := gin.New()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"data": map[string]interface{}{
				"name": "zys",
				"age":  18,
			},
			"message": "success",
		})
	})
	//engine.
	engine.Run(":80")
	//	fmt.Println(engine)
}
