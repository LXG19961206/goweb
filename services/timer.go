package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimerController(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "timer.html", gin.H{
		"time":  time.Now(),
		"title": "测试",
	})

}
