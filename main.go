package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var products = []Product{}

//Product defines the ID and Title testing
type Product struct {
	ID    string
	Title string
}

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/products", getProductsHandler).Methods("GET")
	routes.HandleFunc("/products", postProductsHandler).Methods("POST")
	routes.HandleFunc("/products", deleteProductsHandler).Methods("DELETE")

	http.Handle("/", routes)

	http.ListenAndServe(":8080", nil)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(products)
	w.Write(j)
}

func postProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p Product
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &p)

	products = append(products, p)

	j, _ := json.Marshal(p)
	w.Write(j)
}

func deleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	deleteProduct(id)

	j, _ := json.Marshal(products)
	w.Write(j)
}

func deleteProduct(id string) {
	i := 0
	for _, x := range products {
		if id == x.ID {
			copy(products[i:], products[i+1:])
			products[len(products)-1] = Product{}
			products = products[:len(products)-1]
			break
		}
		i++
	}
}
