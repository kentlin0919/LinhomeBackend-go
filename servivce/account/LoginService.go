package servivce

import (
	pojo "linhomebackend/pojo/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.Account{}

func Postlogin(c *gin.Context) {
	// 解析 JSON 主體到結構體中
	var user pojo.Account
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "kk" {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, "NO")
	}
}
