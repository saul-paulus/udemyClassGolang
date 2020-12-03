package app

import (
	"fmt"
	"log"
	"net/http"
	"udemyApi/handler"

	"github.com/gorilla/mux"
)

func StartServer() {
	// router := http.NewServeMux()
	router := mux.NewRouter()

	fmt.Println("Server di 8082")

	router.HandleFunc("/home", handler.Home).Methods(http.MethodGet)
	router.HandleFunc("/customer", handler.GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", handler.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customers_id:[0-9]+}", handler.GetCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8082", router))
}
