package dto

type ProductRequest struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Stock int `json:"stock"`
}