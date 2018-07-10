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
	db := xorm.DB().DB

	err:= db.Ping()
	if err != nil{
		fmt.Println("ping error:", err)
		return
	}
	fmt.Println("Ping Successful")

	var ver string
	err= db.QueryRow("select version()").Scan(&ver)
	if err != nil{
		fmt.Println("version error:", err)
		return
	}
	fmt.Println("version:", ver)

	c := customer{Name: "xorm insert 1"}
	i, err := xorm.Insert(c)
	if err != nil {
		fmt.Println("xorm Insert error:", err)
		return
	}
	fmt.Println("xorm Insert completed:", i)

	_, err= db.Exec(`insert into customer(name) values('manual insert 1') `)
	if err != nil {
		fmt.Println("manual Insert error:", err)
		return
	}
	fmt.Println("manual Insert completed:")
}
