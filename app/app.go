package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	// router := http.NewServeMux()
	router := mux.NewRouter()

	fmt.Println("Server di 8082")

	router.HandleFunc("/home", Home).Methods(http.MethodGet)
	router.HandleFunc("/customer", GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customers_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8082", router))
}
