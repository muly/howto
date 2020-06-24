package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	proto "github.com/muly/howto/golang/web/grpc/uid/proto"
)

type grpcServer struct{}

func (grpcServer) Generate(ctx context.Context, input *proto.Input) (*empty.Empty, error) {
	genTime := time.Now()
	uids := genUniqueIDs(genTime, input.Qty)
	for _, uid := range uids {
		err := dbInsert(uid)
		if err != nil {
			return nil, err
		}
	}
	for _, uid := range uids {
		fmt.Println(uid)
	}
	return &empty.Empty{}, nil
}
