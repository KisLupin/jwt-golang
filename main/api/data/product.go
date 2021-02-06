package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required,description"`
	Price       float32 `json:"price" validate:"gt=0"`
	CreateOn    string  `json:"create_on"`
	UpdateOn    string  `json:"update_on"`
	DeleteOn    string  `json:"_"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("description", validateDescription)
	return validate.Struct(p)
}

func validateDescription(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.Id = genNextId()
	productList = append(productList, p)
}

func genNextId() int {
	lp := productList[len(productList)-1]
	return lp.Id + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.Id = id
	productList[pos] = p
	return nil
}

var ErrorProductNotFound = fmt.Errorf("prodcut not found")

func findProduct(id int) (*Product, int, error) {
	for k, v := range productList {
		if v.Id == id {
			return v, k, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

var productList = []*Product{
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
