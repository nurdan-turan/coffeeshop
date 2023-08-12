package handlers

import (
	. "coffeeshop/models"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var productStore = make(map[int]Product)
var id int = 0

//HTTP GET - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}
	data, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	w.Write(data)
}

//HTTP GET - /api/products/{id}
func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	for _, product = range productStore {
		if product.Id == id {
			break
		}
	}

	data, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	w.Write(data)

}

//HTTP POST - /api/products
func PostProductsHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	id++
	product.Id = id
	productStore[product.Id] = product

	data, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	w.Write(data)
}

//HTTP PUT - /api/products/{id}
func PutProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	var productToUpdate Product
	err = json.NewDecoder(r.Body).Decode(&productToUpdate)
	if err != nil {
		panic(err)
	}

	if _, ok := productStore[id]; ok {
		productToUpdate.Id = id
		productStore[id] = productToUpdate
	} else {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent) //204
}

//HTTP DELETE - /api/products/{id}
func DeleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	if _, ok := productStore[id]; ok {
		delete(productStore, id)
	} else {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent) //204
}
