package main

import (
	"fmt"
	"os"
	"runtime"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/lhboy1984/rpc-bench/thrift/mathservice"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(8))
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket("127.0.0.1:9090")
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := mathservice.NewMathServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in 127.0.0.1:9090")
	server.Serve()
}

type RpcServiceImpl struct {
}

func (s *RpcServiceImpl) Add(A int32, B int32) (r int32, err error) {
	return A + B, nil
}
