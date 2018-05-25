package main

import (
	"context"
	"fmt"
	pb "github.com/muly/howto/golang/web/grpc/echo-example/pb"
	"google.golang.org/grpc"
	"io"
	"strconv"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	//
	/*text := &pb.Input{Text: "hello, world"}
	resp, err := client.EchoEcho(context.Background(), text)
	if err != nil {
		fmt.Println("client call error: ",err)
		return
	}
	fmt.Println("response is:",resp.Text)*/

	// stream output
	/*
		text := &pb.Input{Text: "hello, stream"}
		resp, err := client.EchoEchoOutputStream(context.Background(), text)
		if err != nil {
			fmt.Println("client call error: ",err)
			return
		}

		out := &pb.OutputStream{}

		for i:=0; ;i++{
			out, err= resp.Recv()
			if err == io.EOF{
				break
			}
			if err != nil {
				fmt.Println("client Recv() error: ",err)
				return
			}
			fmt.Printf("response %v is:%v\n",i, out)
		}
	*/

	// stream input
	/*
		streamCLient, err:= client.EchoEchoInputStream(context.Background())
		if err != nil{
			fmt.Println(err)
		}

		for i:= 0; i<10; i++{
			in:= &pb.Input{"input stream-"+strconv.Itoa(i)}
			fmt.Println("sending ", in.Text)
			err := streamCLient.Send(in)
			if err != nil{
				fmt.Println(err)
			}
		}

		resp, err := streamCLient.CloseAndRecv()
		if err != nil {
			fmt.Println("client stream close error:", err)
			return
		}

		fmt.Println("server response: ", resp.Text)
	*/

	//// Bi stream
	BiStream, err := client.EchoEchoBiStream(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan byte)

	// receive from server
	go func() {
		for {
			o, err := BiStream.Recv()
			if err == io.EOF {
				close(ch)
				break
			}
			fmt.Println("receiving", o.Text)
		}
	}()

	// send to server
	for i := 0; i < 10; i++ {
		in := &pb.OutputStream{"input stream-" + strconv.Itoa(i)}
		fmt.Println("sending ", in.Text)
		err := BiStream.Send(in)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = BiStream.CloseSend()
	if err != nil {
		fmt.Println("client stream close error:", err)
		return
	}

	<-ch // or else sleep for some time to give the go routine to complete execution

}
