// Note: make sure pg is already installed.
// eventhough the table doesn;t exist, the Sync2() will create the table and the columns.
// if the table already exists but not any columns, then the new columns are created.

package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type customer struct {
	Name string `xorm:"text name2"`
	Address string `xorm:"text address"`
}

func initORM() *xorm.Engine {

	ds := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

	x, err := xorm.NewEngine("postgres", ds)
	if err != nil {
		panic(err)
	}
	return x

}

func main() {
	xorm := initORM()

	xorm.Sync2(&customer{})

	c := customer{Name: "xorm insert 3"}
	i, err := xorm.Insert(c)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}
	fmt.Println("xorm Insert completed:", i)

}
