package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8080") // assuming server on localhost with this port
	if err != nil {
		fmt.Println(err)
		return
	}

	//Note: match the data type with the server function's parameters
	err = c.Call("MyServer.Hello", 0, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server Hello function executed")

	out := ""
	err = c.Call("MyServer.Echo", "Good morning", &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server Echo function executed. received:", out)
}
