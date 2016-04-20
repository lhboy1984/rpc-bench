package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/lhboy1984/rpc-bench/grpc/mathservice"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client mathservice.MathServiceClient
}

func (cli *GrpcClient) Close() {
	if cli.conn != nil {
		cli.conn.Close()
		cli.conn = nil
	}
}

func (cli *GrpcClient) Send() {
	cli.client.Add(context.Background(), &mathservice.AddRequest{A: 1, B: 2})
}

func NewGrpcClient(addr string) (cli *GrpcClient, err error) {
	cli = &GrpcClient{}

	cli.conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		return
	}

	cli.client = mathservice.NewMathServiceClient(cli.conn)
	return
}
