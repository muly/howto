// to demonstrate how to connect to a database, run query, convert to json and print json
// uses a struct to store the row data before converting to json
// uses postgres as the sql database in this example

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type address struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	data := []address{}
	rows, err := db.Query("SELECT id, name, address FROM public.address")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name, addr string
		var id int

		err = rows.Scan(&id, &name, &addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		data = append(data, address{Id: id, Name: name, Address: addr})
	}

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j)) // [{"id":12345,"name":"C","address":"USA"}]
}

// database scripts
/*
CREATE TABLE public.address
(
    name character varying(100) COLLATE pg_catalog."default",
    address character varying(100) COLLATE pg_catalog."default",
    id bigint
);

insert into public.address(id, name, address) values(12345, 'C', 'USA');

*/
