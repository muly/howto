// Note: inspired from the example https://github.com/kmanley/golang-rpc-example
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type MyServer struct{}

// properties of the correct rpc function:
// all rpc functions must have 3 input params: receiver, *args, *reply
// 	first input parameter is the receiver itself
// 	third input parameter must be pointer.
// and one output parameter, and its return type must be error

func (MyServer) Hello(args int64, reply *int64) error {
	fmt.Println("Hello World")
	return nil
}

func (MyServer) Echo(args string, reply *string) error {
	fmt.Println("received & echoed:", args)
	*reply = args
	return nil
}

func main() {
	err := rpc.Register(new(MyServer))
	if err != nil {
		fmt.Println(err)
		return
	}
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
