package main

import (
	"github.com/KisLupin/jwt-golang/main/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	l.Println("Server on port 9090")

	hh := handler.NewProducts(l)
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", hh.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", hh.UpdateProduct)
	putRouter.Use(hh.MiddlewareValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", hh.AddProduct)
	postRouter.Use(hh.MiddlewareValidation)

	configureServer(sm, l)
}

func configureServer(sm *mux.Router, l *log.Logger)  {
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Fatal(s.ListenAndServe())
}
