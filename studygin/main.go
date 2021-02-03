package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
)

func main()  {
	fmt.Println("GIN Framework  Study")

	//gin.Default()
	gin.SetMode(config.AppMode)
	engine := gin.New()
}
