package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/estebanmpa/go-pure-rest-api/database"
	"github.com/estebanmpa/go-pure-rest-api/utils"
	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(router *mux.Router) {
	router.HandleFunc("/customer", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomerById).Methods("GET")
	router.HandleFunc("/customer", saveCustomer).Methods("POST", "PUT")
	router.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	cus := database.RetrieveCustomers()
	json.NewEncoder(w).Encode(cus)
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	params := mux.Vars(r)
	id := utils.ToInt(params["id"])
	cus := database.RetrieveCustomerById(id)
	json.NewEncoder(w).Encode(cus)
}

func saveCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer database.Customer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Invalid")
	}
	json.Unmarshal(body, &newCustomer)

	if newCustomer.ID != 0 {
		newCustomer = database.UpdateCustomer(newCustomer)
	} else {
		newCustomer = database.CreateCustomer(newCustomer)
	}
	json.NewEncoder(w).Encode(newCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	params := mux.Vars(r)
	id := utils.ToInt(params["id"])
	database.DeleteCustomer(id)
}
