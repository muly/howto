package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	_ "gitscm.cisco.com/ccdev/go-common/structs"
)

func connect() *gocql.Session {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "tweets"
	cluster.Port = 9042
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = newAuthenticator()

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}

func newAuthenticator() gocql.Authenticator {
	username := ""
	password := ""

	if username != "" && password != "" {
		return gocql.PasswordAuthenticator{
			Username: username,
			Password: password,
		}
	}
	panic("No DB credentials provided; connection will be attempted without authentication")
}

func main() {
	session := connect()
	defer session.Close()

	stmt, names := qb.Select("tweet_summary").
		Columns("year", "user_id", "total").
		Where(qb.Eq("year"), qb.Eq("user_id")).
		AllowFiltering().
		ToCql()

	bv := []interface{}{"2019", "1234567890"}

	log.Println(stmt)
	log.Println(names)
	log.Println(bv)

	q := gocqlx.Query(session.Query(stmt), names).Bind(bv...)
	defer q.Release()

	var rs []struct {
		Year   string `structs:"year"`
		UserID string `structs:"user_id"`
		Total  string `structs:"total"`
	}
	if err := q.Select(&rs); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", rs)
}
