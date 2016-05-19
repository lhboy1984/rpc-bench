package main

import (
	"fmt"
	"io"
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

type GrpcStreamClient struct {
	conn   *grpc.ClientConn
	client mathservice.MathServiceClient
	stream mathservice.MathService_AddByStreamClient
	waitc  chan struct{}
}

func (cli *GrpcStreamClient) Close() {
	cli.stream.CloseSend()
	<-cli.waitc
}

func (cli *GrpcStreamClient) Send() {
	cli.stream.Send(&mathservice.AddRequest{A: 1, B: 2})
}

func NewGrpcStreamClient(addr string) (cli *GrpcStreamClient, err error) {
	cli = &GrpcStreamClient{}
	cli.conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	cli.client = mathservice.NewMathServiceClient(cli.conn)

	cli.stream, err = cli.client.AddByStream(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	cli.waitc = make(chan struct{})
	go func() {
		for {
			_, err := cli.stream.Recv()
			if err == io.EOF {
				close(cli.waitc)
				return
			} else if err != nil {
				fmt.Println(err)
			}
		}
	}()

	return
}
