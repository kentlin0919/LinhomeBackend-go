package servivce

import (
	pojo "linhomebackend/pojo/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.Account{}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Su")
}

func Postlogin(c *gin.Context) {
	user := pojo.Account{}

	userList = append(userList, user)
	c.JSON(http.StatusOK, "Su")

}
