package src

import (
	servivce "linhomebackend/servivce/account"

	"github.com/gin-gonic/gin"
)

// @Summary 登入
// @Description 使用者登入功能
// @Tags User
// @Accept json
// @Produce json
// @Param Login body object true "Body" {"username":"string","password":"string"}
// @Success 200 {string} string "登入成功"
// @Failure 400 {string} string "請求失敗"
// @Router /user/Login [post]
func PostAccount(r *gin.RouterGroup) {
	user := r.Group("/user")
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
