package model

type Order struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Count int    `json:"count"`
	Sale  int    `json:"sale"`
}
