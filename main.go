package main

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

	"hankliu.com.cn/go-my-blog/routes"
	"hankliu.com.cn/go-my-blog/share/extension"
)

func main() {
	fmt.Println("hello world")

	router := gin.Default()

	extension.Init()
	routes.InitialRoutes(router)
	router.Run()
}
