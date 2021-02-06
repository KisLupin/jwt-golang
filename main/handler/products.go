package handler

import (
	"github.com/KisLupin/jwt-golang/main/api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)
		regex := regexp.MustCompile(`/([0-9])`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("no more than one id")
			http.Error(rw, "invalid url", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("no more than one capture")
			http.Error(rw, "invalid url", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("invalid id cant not convert ", idString)
			http.Error(rw, "invalid url", http.StatusBadRequest)
			return
		}
		p.l.Println("got id: ", id)
		p.updateProduct(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p Products) getProducts(rw http.ResponseWriter) {
	p.l.Println("Get data products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadRequest)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod : %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "product not found", http.StatusBadRequest)
		return
	}
}
