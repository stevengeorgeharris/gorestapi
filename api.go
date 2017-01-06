package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Business represents accepted values
type Business struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    string   `json:"type,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// Address represents accepted values
type Address struct {
	Street   string `json:"street,omitempty"`
	City     string `json:"city,omitempty"`
	Postcode string `json:"postcode,omitempty"`
}

// businesses to hold mock data
var businesses []Business

// GetBusinessEndpoint fetches business
func GetBusinessEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range businesses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Business{})
}

// GetBusinessesEndpoint fetches and returns all businesses
func GetBusinessesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(businesses)
}

// CreateBusinessEndpoint adds new business
func CreateBusinessEndpoint(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var business Business
	err := json.NewDecoder(req.Body).Decode(&business)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	businesses = append(businesses, business)
	json.NewEncoder(w).Encode(business)
	fmt.Println(business)
}

func main() {
	router := mux.NewRouter()
	businesses = append(businesses, Business{ID: "1", Name: "Flour Bakery", Address: &Address{Street: "123 Church Street", City: "London"}})
	businesses = append(businesses, Business{ID: "2", Name: "Milk and Honey", Address: &Address{Street: "24 Shoreditch High Street", City: "London"}})
	router.HandleFunc("/businesses", GetBusinessesEndpoint).Methods("GET")
	router.HandleFunc("/businesses", CreateBusinessEndpoint).Methods("POST")
	router.HandleFunc("/business/{id}", GetBusinessEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}
