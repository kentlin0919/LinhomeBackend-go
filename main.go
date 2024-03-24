package main

import (
	_ "linhomebackend/docs"

	config "linhomebackend/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Linhome
// @version 1.0
// @description swagger test example
// @schemes http
// @host 192.168.0.158:9000
// @contact.name Skyler
// @BasePath /v1
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")
	config.PostAccount(v1)

	router.Run(":9000")
}
