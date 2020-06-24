package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	uid "github.com/muly/howto/golang/web/grpc/uid/proto"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:5050", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := uid.NewUidClient(conn)

	ctx := context.TODO()

	a := uid.Input{Qty: 4096}
	_, err = c.Generate(ctx, &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("processed")
}
