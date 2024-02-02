package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"main.go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		p.getProducts(rw, request)
		return
	case http.MethodPost:
		p.addProduct(rw, request)
		return
	case http.MethodPut:
		rx := regexp.MustCompile(`/([0-9]+)`)
		g := rx.FindAllStringSubmatch(request.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number")
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw, request)
		return
	default:
		// catch all
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product body")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}
	data.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product body")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}
