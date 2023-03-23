package main

import (
	"fmt"
	"log"

	"net/http"

	controllers "Latihan2/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	router.HandleFunc("/transactions", controllers.InsertTransaction).Methods("POST")
	router.HandleFunc("/transactions/{transactions_id}", controllers.DeleteTransactions).Methods("DELETE")

	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.InsertProduct).Methods("POST")
	router.HandleFunc("/products/{products_id}", controllers.DeleteProducts).Methods("DELETE")

	router.HandleFunc("/users/{user_id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/transactions/{transactions_id}", controllers.UpdateTransactions).Methods("PUT")
	router.HandleFunc("/products/{products_id}", controllers.UpdateProduct).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
