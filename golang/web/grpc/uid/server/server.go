package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/globalsign/mgo"
	proto "github.com/muly/howto/golang/web/grpc/uid/proto"
)

const (
	mongoHostUrl = "mongodb://uid_mongodb_1:27017"
	mongoDb      = "tasks"
	mongoTable   = "uid"
	// Note: use localhost instead of uid_mongodb_1 when using a mongodb installed on localhost.
	// in this case mongodb is used from the docker image run using docker-compose
)

func Init() uid {
	// db connection
	session, err := mgo.Dial(mongoHostUrl)
	if err != nil {
		log.Fatalln("mgo.Dial error :", err)
		os.Exit(1)
	}

	return uid{
		session: session,
		db:      session.DB(mongoDb).C(mongoTable),
	}
}

func main() {
	state = Init()
	defer state.session.Close()

	log.Printf("executing server...")

	myServer := grpcServer{}
	s := grpc.NewServer()
	proto.RegisterUidServer(s, myServer)

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln("listener error :", err)
	}

	log.Fatalln(s.Serve(listener))
}
