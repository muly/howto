package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("error dial: ", err)
	}
	blogs := session.DB("test").C("testc")

	type Blog struct {
		ID   bson.ObjectId `bson:"_id,omitempty"`
		Name string
		Tags []string
	}

	m := Blog{Name: "pb development using golang", Tags: []string{"golang", "pb"}}

	err = blogs.Insert(&m)
	if err != nil {
		fmt.Println("error insert: ", err)
	}

	ms := []Blog{}

	err = blogs.Find(nil).All(&ms)
	if err != nil {
		fmt.Println("error find all: ", err)
	}
	fmt.Println(ms)

}
