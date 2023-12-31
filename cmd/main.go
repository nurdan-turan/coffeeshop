package main

import (
	. "coffeeshop/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")

	r := mux.NewRouter()

	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductByIdHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductsHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductsHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductsHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
	log.Println("Stopping the HTTP server on port 8080")
}

/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Nurdan!"))
		})
	func aboutHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)

*/
