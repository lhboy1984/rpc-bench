package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)
import "github.com/lhboy1984/rpc-bench/grpc/mathservice"

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	rpcserv := grpc.NewServer()
	mathservice.RegisterMathServiceServer(rpcserv, &mathService{})
	rpcserv.Serve(lis)
}

type mathService struct {
}

func (s *mathService) Add(ctx context.Context, req *mathservice.AddRequest) (*mathservice.AddReply, error) {
	return &mathservice.AddReply{X: req.A + req.B}, nil
}
