package config

import (
	"fmt"
	pojo1 "linhomebackend/pojo/account"
	servivce "linhomebackend/servivce/account"

	"github.com/gin-gonic/gin"
)

var pojoaccount = pojo1.Account{}

type acc struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

// @Summary 登入
// @Description 使用者登入功能
// @Tags User
// @Accept json
// @Produce json
// @Param  body body pojo1.Account{} true "使用者登入資訊" 
// @Success 200 {string} string "登入成功"
// @Failure 400 {string} string "請求失敗"
// @Router /user/Login [post]
func PostAccount(r *gin.RouterGroup) {
	user := r.Group("/user")
	fmt.Println(pojoaccount)
	user.POST("/Login", servivce.Postlogin)
}

// @Summary 登入
// @Description 使用者登入功能
// @Tags User
// @Accept json
// @Produce json
// @Param username query string true "使用者名稱"
// @Param password query string true "密碼"
// @Success 200 {string} string "登入成功"
// @Failure 400 {string} string "請求失敗"
// @Router /user/Login1 [post]
func PostRegister(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.POST("/Login1", servivce.Postlogin)
}
