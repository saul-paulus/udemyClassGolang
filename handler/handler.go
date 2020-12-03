package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Haiii saul")
}

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{Name: "saul", City: "Ambon", Zipcode: "121311"},
		{Name: "El", City: "Ambon", Zipcode: "12121111"},
	}
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customer)

	if r.Header.Get("Content-Type") == "application/xml" {
		// ubah ke xml
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
