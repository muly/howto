// Note: make sure pg is already installed and a table with the name customer (fields: name) already exists.
//

package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)


func initORM() *xorm.Engine {

	ds := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

	x, err := xorm.NewEngine("postgres", ds)
	if err != nil {
		panic(err)
	}

	x.ShowSQL(true)

	return x

}


type Customer struct {
	Id   int    `xorm:"int pk id"`
	Name string `xorm:"text name2"`
	//Address
}
type Contact struct {
	CustomerId int    `xorm:"int customerid"`
	Type       string `xorm:"text type"`
	Contact    string `xorm:"text contact"`
}
type customercontact struct{
	Customer `xorm:"extends"`
	Contact `xorm:"extends"`
}
func (customercontact) TableName() string {
	return "customer"
}

func main() {
	xorm := initORM()

	xorm.Sync2(&Customer{})
	xorm.Sync2(&Contact{})

	prepare(xorm)

	c := []customercontact{}
	err := xorm.Select("customer.*,contact.*").
		Join("LEFT", "contact", "customer.id=contact.customerid").
		Find(&c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)

}

func prepare(xorm *xorm.Engine) {

	c1 := Customer{Id: 1, Name: "c1"}
	c2 := Customer{Id: 2, Name: "c2"}
	c := []Customer{c1, c2}

	_, err := xorm.Insert(c)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}

	c1a := Contact{CustomerId: 1, Type: "mobile", Contact: "1234567890"}
	c1b := Contact{CustomerId: 1, Type: "work", Contact: "9999999999"}
	c2a := Contact{CustomerId: 2, Type: "email", Contact: "abc@xyz"}
	ct := []Contact{c1a, c1b, c2a}
	_, err = xorm.Insert(ct)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}
}
