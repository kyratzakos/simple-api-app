package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-api-app/routes"

	"github.com/gorilla/mux"
)

func main() {
	routes.InitDB()

	router := mux.NewRouter()

	router.HandleFunc("/products", routes.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", routes.GetProduct).Methods("GET")
	router.HandleFunc("/products", routes.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", routes.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")

	port := os.Getenv("PORT")
	fmt.Printf("Server running at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
