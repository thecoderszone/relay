package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	// HOST is the HTTP server host
	HOST = "localhost"
	// PORT is the HTTP server port
	PORT = 8080
)

func main() {
	MigrateDB()

	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/relays", CreateRelay).Methods("POST")
	router.HandleFunc("/relays/{id}", GetRelay).Methods("GET")
	router.HandleFunc("/relays/{id}", UpdateRelay).Methods("PUT", "PATCH")
	router.HandleFunc("/relays/{id}", DeleteRelay).Methods("DELETE")

	log.Println(fmt.Sprintf("HTTP server started at http://%s:%d", HOST, PORT))
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), router)

	log.Fatalln(err)
}
