package main

import (
	"github.com/KisLupin/jwt-golang/main/handler"
	"log"
	"net/http"
	"os"
)
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	http.ListenAndServe("9090", sm )
}
