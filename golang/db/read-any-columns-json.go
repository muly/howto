// to demonstrate how to read the data from database using a dynamic query, convert to json and print json
// we cannot use a struct to store the intermediate row level data as we dont know the columns that are returned by the dynamic query
// uses postgres as the sql database in this example

// references:
// Working with Unknown Columns: http://go-database-sql.org/varcols.html
// https://kylewbanks.com/blog/query-result-to-map-in-golang

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

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
	err := query()
	if err != nil {
		fmt.Println(err)

	}
}

func query() error {
	rows, _ := db.Query("SELECT id, name, address FROM public.address") // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()

	n := []map[string]interface{}{}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		n = append(n, m)
	}

	j, err := json.Marshal(n)
	if err != nil {
		return err
	}
	fmt.Println(string(j)) // [{"address":"USA","id":12345,"name":"C"}]

	return nil
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
