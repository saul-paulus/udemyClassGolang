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

	router.HandleFunc("/home", handler.Home)
	router.HandleFunc("/customer", handler.GetAllCustomer)
	log.Fatal(http.ListenAndServe(":8082", router))
}
