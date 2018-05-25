package main

import (
	"fmt"
	pb "github.com/muly/howto/golang/web/grpc/echo-example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
)

func main() {

	fmt.Println("listening...")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("could not listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type server struct{}

func (server) EchoEcho(ctx context.Context, i *pb.Input) (*pb.Output, error) {
	o := pb.Output{i.Text}
	return &o, nil
}

func (server) EchoEchoOutputStream(in *pb.Input, s pb.Echo_EchoEchoOutputStreamServer) error {

	for i := 0; i < 10; i++ {
		o := pb.OutputStream{in.Text + "-" + strconv.Itoa(i)}
		err := s.Send(&o)
		if err != nil {
			return err
		}
	}

	return nil
}

func (server) EchoEchoInputStream(s pb.Echo_EchoEchoInputStreamServer) error {
	i := 0
	for ; ; i++ {
		out, err := s.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("server Recv() error: ", err)
			return err
		}
		fmt.Printf("received %v is:%v\n", i, out)
	}

	err := s.SendAndClose(&pb.Output{"received " + strconv.Itoa(i) + "messages"})
	if err != nil {
		fmt.Println("server Recv() error: ", err)
		return err
	}
	return nil
}

func (server) EchoEchoBiStream(s pb.Echo_EchoEchoBiStreamServer) error {
	for {
		o, err := s.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("receiving", o.Text)

		fmt.Println("sending", o.Text)
		err = s.Send(o)
		if err != nil {
			fmt.Printf("error sending message %v: %v\n", o.Text, err)
		}
	}

	return nil
}
