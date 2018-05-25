package main

import (
	"database/sql"
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
	rows, err := db.Query("SELECT id, name, address FROM public.address")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name, address, id string
		err = rows.Scan(&id, &name, &address)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, name, address)
	}
}
