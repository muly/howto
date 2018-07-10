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
type Address struct {
	CustomerId int    `xorm:"int customerid "`
	Line1      string `xorm:"text line1"`
	City       string `xorm:"text city"`
}

type customerAddress struct{
	Customer `xorm:"extends"`
	Address `xorm:"extends"`
}
func (customerAddress) TableName() string {
	return "customer"
}

func main() {
	xorm := initORM()

	xorm.Sync2(&Customer{})
	xorm.Sync2(&Address{})

	prepare(xorm)

	c := []customerAddress{}
	err := xorm.Select("customer.*,address.*").Join("LEFT", "address", "customer.id=address.customerid").
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

	a1 := Address{CustomerId: 1, Line1: "100 N St"}
	_, err = xorm.Insert(a1)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}
}
