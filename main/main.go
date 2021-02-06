package main

import (
	"github.com/KisLupin/jwt-golang/main/handler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	l.Println("Server on port 9090")
	hh := handler.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Fatal(s.ListenAndServe())
}
