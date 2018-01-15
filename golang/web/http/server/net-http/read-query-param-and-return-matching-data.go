// query parameters example along with returning matching records.
// Note: returned data is plain text (not json) for simplicity

// sample data
// id, name, country
// 1, A, usa
// 2, B, india

// input 1: http://localhost:8080/customer
// Output: should have both records A & B

// input 2: http://localhost:8080/customer?id=1
// Output: should have record A

package main

import (
	"fmt"
	"net/http"
)

type customer struct {
	id      string
	name    string
	country string
}

var customerList []customer

func main() {
	loadTestCustomer()

	http.HandleFunc("/customer", customerHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	// read query parameters
	m := r.URL.Query()

	id := ""
	val, ok := m["id"]
	if ok == true {
		id = val[0]
	}
	fmt.Println(id)

	resCustomerList := []customer{}
	for _, c := range customerList {
		if id == "" || c.id == id {
			resCustomerList = append(resCustomerList, c)
		}
	}

	response := fmt.Sprintf("Hello! \nhere is the list of matching customers \n%#v", resCustomerList)
	w.Write([]byte(response))
}

func loadTestCustomer() {
	customerList = append(customerList, customer{id: "1", name: "A", country: "usa"})
	customerList = append(customerList, customer{id: "2", name: "B", country: "india"})
}
