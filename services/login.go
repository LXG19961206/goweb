package services

import (
	"ginWeb/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func LoginController(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}

func AuthController(ctx *gin.Context) {

	var form = &model.UserInfo{}

	_ = ctx.ShouldBindWith(&form, binding.Form)

	var resp = &Resp[*model.UserInfo]{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data: &model.UserInfo{
			Name: form.Name, Password: form.Password,
		},
	}

	ctx.JSON(http.StatusOK, *resp)

}
