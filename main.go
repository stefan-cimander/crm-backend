package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get customers...")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get customer...")
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add customer...")
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update customer...")
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete customer...")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting the CRM backend server on port 3000...")
	http.ListenAndServe(":3000", router)
}
