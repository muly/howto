package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/muly/howto/golang/web/grpc/uid/proto"
)

func main() {
	myServer := grpcServer{}
	s := grpc.NewServer()
	proto.RegisterUidServer(s, myServer)

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln("listener error :", err)
	}

	log.Fatalln(s.Serve(listener))
}
