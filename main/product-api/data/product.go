package data

import "time"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
	CreateOn    string
	UpdateOn    string
	DeleteOn    string
}

func GetProducts() []*Product {
	return product
}

var product = []*Product{
	&Product{
		Id:          1,
		Name:        "Milk",
		Description: "To Drink",
		Price:       30.25,
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		Id:          2,
		Name:        "Coca cola",
		Description: "To Drink",
		Price:       23.25,
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
