package handler

import (
	"log"
	"net/http"
	//"github.com/KisLupin/jwt-golang/main/api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) * Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//lp := data.GetProducts()
}