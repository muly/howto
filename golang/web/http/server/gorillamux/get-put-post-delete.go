package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Customers []customer

func main() {
	//add records to structs
	Customers = append(Customers, customer{1, "Bobby"})
	Customers = append(Customers, customer{2, "Robin"})

	router := mux.NewRouter()
	//customer pb
	router.HandleFunc("/customers", customerHandler).Methods("GET")                    //gets all customers
	router.HandleFunc("/customers/{id}", getCustomerHandler).Methods("GET")            //gets one customer
	router.HandleFunc("/customers/{id}/{name}", editCustomerHandler).Methods("PUT")    //gets one customer
	router.HandleFunc("/customers/{id}/{name}", createCustomerHandler).Methods("POST") //creates new customer
	router.HandleFunc("/customers/{id}", deleteCustomerHandler).Methods("DELETE")      //deletes a customer

	log.Fatal(http.ListenAndServe(":8080", router))
}
func jMarshal(w http.ResponseWriter, r *http.Request, res []customer) {
	jResult, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "marshal Error...:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jResult)
}
func customerHandler(w http.ResponseWriter, r *http.Request) {
	//fetch all customers
	jResult, err := json.Marshal(Customers)
	if err != nil {
		http.Error(w, "marhsal error", http.StatusInternalServerError)
	}
	w.Write(jResult)
}
func getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jResult := []customer{} //need to empty the slice before appending query results
	for _, c := range Customers {
		if strconv.Itoa(c.Id) == vars["id"] {
			jResult = append(jResult, c)
		}
	}
	jMarshal(w, r, jResult) //gives key and values whereas fmt.Fprint gives only values
}
func editCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "type conversion error...")
		return
	}
	for i, j := range Customers {
		if j.Id == nId {
			Customers[i].Name = vars["name"]
		}
	}
	jMarshal(w, r, Customers)

}
func createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "type conversion error", http.StatusInternalServerError)
	}

	for _, i := range Customers {
		if i.Id == nId {
			http.Error(w, "customer already exists", http.StatusConflict)
			return
		}
	}
	var cus customer
	cus.Id = nId
	cus.Name = vars["name"] //did this by split(/) of vars
	Customers = append(Customers, cus)
	jMarshal(w, r, Customers)
}
func deleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "conversion error", http.StatusInternalServerError)
	}
	for i, j := range Customers { //need how this worked ?
		if j.Id == nId {
			Customers = append(Customers[:i], Customers[i+1:]...)
			//delete(Customers,nId)//helpful only if custms is a map of slice
			break
		}
	}
	jMarshal(w, r, Customers)
}
