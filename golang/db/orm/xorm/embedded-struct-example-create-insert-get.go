// example to demonstrate that struct embedding will work with xorm
// solution: the important point is to
// keep your field names exported,
// and the alias for embedded struct is also exported is provided
// and 'extends' xorm tag is used for the embedded structs

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
	Id      int
	Name    string
	Contact `xorm:"extends"` // no aliasing, and the embedded struct is already exported
	T cat `xorm:"extends"` // aliasing the unexported embedded struct
}
type Contact struct {
	Address string
	Email   string
}
type cat struct {
	Category    string
	SubCategory string
}

func main() {
	xorm := initORM()

	// sync/create
	xorm.Sync2(&Customer{})

	//insert
	c := Customer{Name: "A", Contact: Contact{Address: "123 St", Email: "a@abc.com"}, T: cat{Category: "cat1", SubCategory: "sub cat1"}}
	_, err := xorm.Insert(c)
	if err != nil {
		fmt.Println(err)
	}

	//get
	g := []Customer{}
	err = xorm.Find(&g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g)

}

// after the sync operation, table is created with the following schema
//CREATE TABLE public.customer
//(
//id 		integer,
//name 		character varying(255),
//address 	character varying(255),
//email 	character varying(255),
//category 	character varying(255),
//sub_category character varying(255)
//)
