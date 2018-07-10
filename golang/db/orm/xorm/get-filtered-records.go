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
	Id   int
	Name string
}

type CustomerList []Customer

func (d *CustomerList) GetAll() error {
	xorm := initORM()

	err := xorm.Find(d)
	if err != nil {
		return err
	}

	return nil
}
func get(d interface{}) error {
	xorm := initORM()

	err := xorm.
		//Where("name2 = ?", "A").
		Where("id >= ?", 2).Where("id <= ?", 4).
		Find(d)
	if err != nil {
		return err
	}
	return nil
}
func (d *CustomerList) GetFiltered() error {
	return get(d)
}


func (d *CustomerList) GetFilteredBetween() error {
	xorm := initORM()

	err := xorm.
		//Where("name2 = ?", "A").
		Where("id between ? and ?", 2, 4).
		Find(d)
	if err != nil {
		return err
	}
	return nil
}

func (d *CustomerList) GetFilteredRaw() error {
	xorm := initORM()

	res, err := xorm.Query("select id, name2 from customer_table where name2 = ?", "A")
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func main() {
	xorm := initORM()
	xorm.Sync2(&Customer{})

	c := CustomerList{}
	err := c.GetFiltered()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}
