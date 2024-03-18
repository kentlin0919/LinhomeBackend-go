package main

import (
	"linhomebackend/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	src.AddAcoutrouter(v1)

	router.Run(":8080")
}
