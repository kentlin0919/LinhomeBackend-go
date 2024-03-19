package src

import (
	account "linhomebackend/src/account"

	"github.com/gin-gonic/gin"
)

func AccountRouter(r *gin.RouterGroup) {
	account.PostAccount(r)
	account.PostRegister(r)
}
