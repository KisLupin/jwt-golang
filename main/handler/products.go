package handler

import (
	"context"
	"fmt"
	"github.com/KisLupin/jwt-golang/main/model/data"
	"github.com/gorilla/mux"
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
		return
	}
	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
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
		p.UpdateProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.l.Println("Get data products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadRequest)
	}
}

func (p *Products) AddProduct(_ http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod : %#v", prod)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable concert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "product not found", http.StatusBadRequest)
		return
	}
}

type KeyProduct struct {

}

func (p Products) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("unable to unmarshal to json")
			http.Error(rw, "unable to unmarshal to json", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("error invalid data ", err)
			http.Error(rw, fmt.Sprintf("error invalid data %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
