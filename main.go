package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/estebanmpa/go-pure-rest-api/config"
	"github.com/estebanmpa/go-pure-rest-api/database"
	"github.com/estebanmpa/go-pure-rest-api/endpoints"
)

func main() {
	registerRoutes()
	testDataBase()
	startServer()
}

func registerRoutes() {
	router := mux.NewRouter()
	endpoints.RegisterCustomerRoutes(router)
	endpoints.RegisterProductRoutes(router)
	http.Handle("/", router)
}

func startServer() {
	fmt.Printf("Starting server on port " + config.GetPort() + "...\n")
	err := http.ListenAndServe(config.GetPort(), nil)
	if err != nil {
		panic(err.Error())
	}
}

func testDataBase() {
	db := database.GetConnection()
	err := db.Ping()
	if err != nil {
		panic("Database Connection Issue")
	} else {
		fmt.Println("Connected to database")
	}
}
