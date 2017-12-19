// data exchange format: json
// just using net/http package
// persist data just in memory
// assume below data struct (use slice of this struct to store multiple records)
// type customer struct{Id int, Name string, Phone string, City string, Gender string}
// below endpoints
// GET
// /customer : get all customers
// /customer/1 : get customer with id=1
// /customer?name=abc : get all customers with name = abc
// /customer?city=hyderabad : get all customers with city = hyderabad, likewise search by phonenumber, or gender
// /customer?city=hyderabad&gender=male : like wise in any combination
// PUT
// /customer/1 : modify the customer 1 record with the new data given in the body of the request
// DELETE
// /customer/1 : delete customer 1 record
// POST
// /customer : add new customer record with the given data in the body of the request

//r.url.Parse() key value

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type customer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	City   string `json:"city"`
}

var c map[string]customer // map of id to the record

func main() {

	c = map[string]customer{}

	c["1"] = customer{1, "Bobby", "Male", "1234567890", "Hyderabad"}
	c["2"] = customer{2, "Robin", "Male", "1234509876", "Bangalore"}
	c["3"] = customer{3, "Bobby", "Male", "0987612345", "Hyderabad"}

	http.HandleFunc("/customers", customers)
	http.HandleFunc("/customer/", custom)
	http.HandleFunc("/customer", rQuery)
	http.ListenAndServe(":8080", nil)
}

func customers(w http.ResponseWriter, r *http.Request) {
	c, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(c))
}

func custom(w http.ResponseWriter, r *http.Request) {
	q := strings.Split(r.URL.Path, "/")[2]

	result := []customer{}
	for _, i := range c {
		if strconv.Itoa(i.Id) == q || strings.ToLower(i.Name) == strings.ToLower(q) {
			result = append(result, i)
		}
	}
	fmt.Fprint(w, result)
	return
	//w.Write(result)
}

func jMarshal(w http.ResponseWriter, r *http.Request, res []customer) {
	jResult, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Marshal Error...:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jResult)
}

func rQuery(w http.ResponseWriter, r *http.Request) {
	result := []customer{}
	q := r.URL.Query()
	qId := q["id"][0]
	cus := customer{}

	switch r.Method {
	case "GET":
		if len(q) == 0 {
			jResult, err := json.Marshal(c)
			if err != nil {
				http.Error(w, "Marshal Error...:"+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jResult)
			return
		}
		qId, exists := q["id"]
		if !exists {
			for _, j := range c {
				if strconv.Itoa(j.Id) == qId[0] {
					result = append(result, j)
				}
			}
			jMarshal(w, r, result)
			return
		}
		qName, exists := q["name"]
		if !exists {
			for _, j := range c {
				if j.Name == qName[0] {
					result = append(result, j)
				}
			}
			jMarshal(w, r, result)
			return
		} else {
			for _, j := range c {
				if j.Name == qName[0] && strconv.Itoa(j.Id) == qId[0] {
					result = append(result, j)
				}
			}
			jMarshal(w, r, result)
			return
		}
	case "PUT":
		for _, i := range c {
			if strconv.Itoa(i.Id) == qId {
				cus.Id, _ = strconv.Atoi(q.Get("id"))
				cus.Name = q.Get("name")
				cus.Gender = q.Get("gender")
				cus.Phone = q.Get("phone")
				cus.City = q.Get("city")
				c[qId] = customer{cus.Id, cus.Name, cus.Gender, cus.Phone, cus.City}
			}
		}
		fmt.Fprint(w, "Customer updated...")
		return
	case "POST":
		for _, i := range c {
			if strconv.Itoa(i.Id) != qId {
				cus.Id, _ = strconv.Atoi(q.Get("id"))
				cus.Name = q.Get("name")
				cus.Gender = q.Get("gender")
				cus.Phone = q.Get("phone")
				cus.City = q.Get("city")
				c[qId] = customer{cus.Id, cus.Name, cus.Gender, cus.Phone, cus.City}
				fmt.Fprintf(w, "Customer posted...")
				break
			} else {
				http.Error(w, "Customer already exists...", http.StatusConflict)
				return
			}
		}
	case "DELETE":
		for _, i := range c {
			if strconv.Itoa(i.Id) == qId {
				delete(c, qId)
				break
			} else {
				http.Error(w, "Customer not found...", http.StatusNotFound)
				return
			}
		}
		fmt.Fprintf(w, "Customer found and deleted...")
	}

}
