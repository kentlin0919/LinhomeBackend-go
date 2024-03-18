package src

import (
	"linhomebackend/servivce"

	"github.com/gin-gonic/gin"
)

func AddAcoutrouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/", servivce.Login)
	user.POST("/", servivce.Postlogin)
}
