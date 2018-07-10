// Mysql Go example: not tested
package main

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var dbConn *gorm.DB

var (
	user     string = "root"
	secret   string = "password"
	dbip     string = "localhost"
	dbport   string = "3306"
	dbschema string = "dev"
)

func init() {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(user, ":", secret, "@tcp(", dbip, ":", dbport, ")/", dbschema))
	log.Println("MySQL Database Connection String :", buffer.String())
	dbURL := buffer.String()

	var err error
	dbConn, err = gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	dbConn.DB()
	err = dbConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(20)
	dbConn.SingularTable(true)
	dbConn.LogMode(true)
	return

}

type Customer struct {
	ID     int    `gorm:"column:ID"  json:"ID"`
	Name   string `gorm:"column:NAME"  json:"name"`
	Email  string `gorm:"column:EMAIL"  json:"email"`
	Status string `gorm:"column:STATUS"  json:"status"`
}

func main() {

	fmt.Println("Hello, playground")

	cust := Customer{Name: "test 1", Email: "test email 1", Status: "active"}
	dbConn.Save(&cust)

	c := []Customer{}
	dbConn.Table("customer").Scan(&c)
	fmt.Println("Active Customers list: ", c)

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


*/
