package services

import "github.com/gin-gonic/gin"

func HelloController(ctx *gin.Context) {

	ctx.String(200, "hello world")

}
