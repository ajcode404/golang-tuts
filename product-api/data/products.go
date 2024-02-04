package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = ""
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	up, _, err := getProductWithId(id)
	if err != nil {
		return err
	}
	up.Name = p.Name
	up.Price = p.Price
	up.Description = p.Description
	up.SKU = p.SKU
	return nil
}

func getProductWithId(id int) (*Product, int, error) {
	for i, d := range productList {
		if d.ID == id {
			return d, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func getProductWithIdMyVersion(id int) (*Product, error) {
	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			return productList[i], nil
		}
	}
	return nil, &MyErrProductNotFound{}
}

type MyErrProductNotFound struct {
}

func (m *MyErrProductNotFound) Error() string {
	return "Product not found"
}

func getNextId() int {
	return productList[len(productList)-1].ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milkey cofee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong cofee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
