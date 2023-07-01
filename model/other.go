package model

type PingQuery struct {
	Name string `form:"name" json:"name"`
	Age  string `form:"age" json:"age"`
	Id   string `form:"id" json:"id"`
}
