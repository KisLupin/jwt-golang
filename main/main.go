package main

import (
	"fmt"
	"github.com/KisLupin/jwt-golang/main/handler"
	"github.com/KisLupin/jwt-golang/main/model/files"
	"github.com/KisLupin/jwt-golang/main/protos"
	"github.com/KisLupin/jwt-golang/main/server"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {
	//api()
	//gRPC()
	mssqlConnection()
}

var serverIP = "14.225.74.30"
var port = 1433
var user = "mobihome_dev"
var password = "mobihome_dev"
var database = "mobihome_dev"

type Order struct {
	OrderGuid uuid.UUID
	OrderNo  int
}

func mssqlConnection() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		serverIP, user, password, port, database)
	// Create connection pool

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	var data []Order
	db.Raw("select * from [Order] o", &data).Scan(&data)
	fmt.Println("data ", data)
	fmt.Printf("Connected!\n")
}

func gRPC() {
	l := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()
	// create an instance of the Currency server
	c := server.NewCurrency(l)
	// register the currency server
	protos.RegisterCurrencyServer(gs, c)
	// register the reflection service which allows clients to determine the methods
	reflection.Register(gs)
	// create a TCP socket for inbound server connections
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		l.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	// listen for requests
	_ = gs.Serve(listen)
}

func api() {
	logger := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	sm := mux.NewRouter()
	basicApi(l, sm)
	executeFile(logger, sm)

	configureServer(sm, l)
}

func executeFile(logger hclog.Logger, sm *mux.Router) {
	stor, err := files.NewLocal(*basePath, 1024*1024*5)
	if err != nil {
		logger.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}
	// create the handlers
	fh := handler.NewFiles(stor, logger)
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(*basePath))),
	)
}

func basicApi(l *log.Logger, sm *mux.Router) {
	hh := handler.NewProducts(l)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", hh.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", hh.UpdateProduct)
	putRouter.Use(hh.MiddlewareValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", hh.AddProduct)
	postRouter.Use(hh.MiddlewareValidation)
}

func configureServer(sm *mux.Router, l *log.Logger) {
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Fatal(s.ListenAndServe())

	//go func() {
	//	l.Info("Starting server on port 9090")
	//
	//	err := s.ListenAndServe()
	//	if err != nil {
	//		l.Error("Error starting server", "error", err)
	//		os.Exit(1)
	//	}
	//}()
	//
	//// trap sigterm or interupt and gracefully shutdown the server
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//signal.Notify(c, os.Kill)
	//
	//// Block until a signal is received.
	//sig := <-c
	//log.Println("Got signal:", sig)
	//
	//// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//s.Shutdown(ctx)
}
