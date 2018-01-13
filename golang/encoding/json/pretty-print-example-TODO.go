package main

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
)

func main() {
	//TODO: define a struct, initialize it and pass to pretty function below

	b, err := prettyjson.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
