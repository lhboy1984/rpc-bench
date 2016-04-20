package main

import (
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/lhboy1984/rpc-bench/thrift/mathservice"
)

type ThriftClient struct {
	trans  thrift.TTransport
	client *mathservice.MathServiceClient
}

func (cli *ThriftClient) Close() {
	if cli.trans != nil {
		cli.trans.Close()
		cli.trans = nil
	}
}

func (cli *ThriftClient) Send() {
	cli.client.Add(1, 2)
}

func NewThriftClient(addr string) (cli *ThriftClient, err error) {
	cli = &ThriftClient{}

	cli.trans, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		return
	}
	cli.trans = thrift.NewTFramedTransport(cli.trans)

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	cli.client = mathservice.NewMathServiceClientFactory(cli.trans, protocolFactory)
	if err := cli.trans.Open(); err != nil {
		cli.Close()
		fmt.Fprintln(os.Stderr, "Error opening socket to ", addr, " ", err)
	}

	return
}
