package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/lib/pq/hstore"
)

var db *sql.DB

func init() {
	var err error
	dburl := os.Getenv("dbserver")
	db, err = sql.Open("postgres", dburl)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//

func main() {
	rows, err := db.Query("select attributes from shard_1.user where id = 123")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		attr, test := hstore.Hstore{}, 10
		err = rows.Scan(&attr)
		if err != nil {
			fmt.Println(err)
			return
		}
		for k, v := range attr.Map {
			fmt.Printf("%v:%v\n", k, v.String)
		}
	}
}
