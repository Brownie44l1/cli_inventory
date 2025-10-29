package model

type Item struct {
	ID int `json:"id"`
	Sku string `json:"sku"`
	Name string `json:"name"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}