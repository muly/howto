package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	proto "github.com/muly/howto/golang/web/grpc/uid/proto"
)

type grpcServer struct {
}

func (grpcServer) Generate(ctx context.Context, input *proto.Input) (*empty.Empty, error) {
	uids := genUniqueIDs(input.Qty)
	for _, uid := range uids {
		err := dbInsert(uid)
		if err != nil {
			log.Fatalln("dbInsert error %#v", err)
			return nil, err
		}
	}

	return &empty.Empty{}, nil
}
