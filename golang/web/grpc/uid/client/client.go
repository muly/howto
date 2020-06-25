package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	uid "github.com/muly/howto/golang/web/grpc/uid/proto"
)

const (
	qty = 100
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:5050", opts...)
	if err != nil {
		log.Println(err)
		return
	}
	c := uid.NewUidClient(conn)

	ctx := context.TODO()

	a := uid.Input{Qty: qty}
	_, err = c.Generate(ctx, &a)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("generated %d uids", qty)
}
