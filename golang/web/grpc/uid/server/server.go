package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/globalsign/mgo"
	proto "github.com/muly/howto/golang/web/grpc/uid/proto"
)

func Init() uid {
	var myStruct uid

	// db connection
	url := "mongodb://localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalln("Dial url error :", err)
		os.Exit(1)
	}
	myStruct.session = session

	// collection object
	myStruct.db = myStruct.session.DB("tasks").C("uid")

	return myStruct
}

func main() {
	state = Init()
	defer state.session.Close()

	myServer := grpcServer{}
	s := grpc.NewServer()
	proto.RegisterUidServer(s, myServer)

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln("listener error :", err)
	}

	log.Fatalln(s.Serve(listener))
}
