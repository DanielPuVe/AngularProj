package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// original main function original 
// func main() {
//    router := mux.NewRouter()
//    log.Fatal(http.ListenAndServe(":8000", router))
//}

type Product struct {
    CustomerId		string   `json:"customerId,omitempty"`
    Item 			string   `json:"item,omitempty"`
    Count  			string   `json:"count,omitempty"`
	Date  			string   `json:"date,omitempty"`
}


var products []Product

func main() {

	router := mux.NewRouter()
	
	products = append(products, Product{CustomerId: "1", Item: "Dani", Count: "7", Date: "02/06/2016"})
	products = append(products, Product{CustomerId: "2", Item: "Dani", Count: "4", Date: "02/06/2016"})
	products = append(products, Product{CustomerId: "3", Item: "Milki", Count: "1", Date: "03/06/2016"})
	products = append(products, Product{CustomerId: "4", Item: "Milki", Count: "2", Date: "04/06/2016"})
	products = append(products, Product{CustomerId: "5", Item: "Milki", Count: "2", Date: "04/06/2016"})
	
    router.HandleFunc("/products", GetProducts).Methods("GET")
    router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
	//http://localhost:8080/products
}

//func GetProducts(w http.ResponseWriter, r *http.Request) {}
func GetProduct(w http.ResponseWriter, r *http.Request) {}
func CreateProduct(w http.ResponseWriter, r *http.Request) {}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {}



func GetProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(products)
}

