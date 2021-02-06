package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreateOn    string  `json:"create_on"`
	UpdateOn    string  `json:"update_on"`
	DeleteOn    string  `json:"_"`
}

type Products []*Product

func (p Products) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return product
}

var product = []*Product{
	{
		Id:          1,
		Name:        "Milk",
		Description: "To Drink",
		Price:       30.25,
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		Id:          2,
		Name:        "Coca cola",
		Description: "To Drink",
		Price:       23.25,
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
