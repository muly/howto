// demonstrate how to connect to mysql database and run a select query

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     string = "root"
	secret   string = "password"
	dbip     string = "localhost"
	dbport   string = "3306"
	dbschema string = "dev"
)

var db *sql.DB

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, secret, dbip, dbport, dbschema)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	rows, err := db.Query("SELECT id, name, email, status FROM customer")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var id, name, email, status string
		err = rows.Scan(&id, &name, &email, &status)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, name, email, status)
	}
}

/*
create database dev;

use dev;

drop table customer;

create table customer(
ID int,
NAME varchar(100),
EMAIL varchar(100),
STATUS varchar(100)
);

INSERT INTO customer values(1, "mysql", "my@sql.com", "active")

*/
