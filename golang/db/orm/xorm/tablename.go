// example to demonstrate how to handle if the Go struct name is not same as the corresponding pg table name.
// the solution is to create a method `TableName()` that returns the corresponding pg table name.

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
}
func (Customer) TableName() string {
	return "customer_table"
}

func main() {
	xorm := initORM()

	xorm.Sync2(&Customer{})

	c := []Customer{}
	err := xorm.Find(&c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}

