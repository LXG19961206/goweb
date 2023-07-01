package services

import (
	"encoding/json"
	"fmt"
	"ginWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingController(ctx *gin.Context) {

	var (
		params           = &model.PingQuery{}
		jsonStr, jsonErr = json.Marshal(params)
		myJson           = `{ "name": "tom", "age": "20", "id": "30" }`
		myVal            model.PingQuery
	)

	usErr := json.Unmarshal([]byte(myJson), &myVal)

	if usErr == nil {
		fmt.Println(myVal)
	}

	if jsonErr == nil {
		fmt.Println(string(jsonStr))
	}

	err := ctx.ShouldBindQuery(params)

	if err != nil {

	}

	var order = model.Order{
		Name:  "手机",
		Price: "100",
		Count: 1000,
	}

	var resp = &Resp[[]model.Order]{
		Code:    200,
		Message: "操作成功",
		Data:    []model.Order{order, order, order},
	}

	ctx.JSON(http.StatusOK, resp)

}
