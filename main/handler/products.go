package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/KisLupin/jwt-golang/main/api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) * Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadRequest)
	}
	_, _ = rw.Write(d)
}