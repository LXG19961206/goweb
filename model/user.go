package model

type UserInfo struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}
