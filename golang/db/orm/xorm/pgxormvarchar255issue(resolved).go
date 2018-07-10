// Note: make sure pg is already installed and a table with the name customer (fields: name) already exists.
//

package xorm

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type customer struct {
	Name string `xorm:"text name"`
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

	err := xorm.Ping()
	if err != nil {
		fmt.Println("ping error:", err)
		return
	}
	fmt.Println("Ping Successful")

	c := customer{Name: "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345"}
	fmt.Println("len=:", len(c.Name))

	err = xorm.Sync2(&c)
	if err != nil {
		fmt.Println("sync error:", err)
		return
	}

	i, err := xorm.Insert(c)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}
	fmt.Println("xorm Insert completed:", i)
}
