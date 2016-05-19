package main

import (
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)
import (
	"github.com/davecheney/profile"
	_ "github.com/icattlecoder/godaemon"
	"github.com/lhboy1984/rpc-bench/grpc/mathservice"
)

func main() {
	defer profile.Start(&profile.Config{CPUProfile: true, MemProfile: true, BlockProfile: true}).Stop()
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

func (s *mathService) AddByStream(stream mathservice.MathService_AddByStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		} else {
			stream.Send(&mathservice.AddReply{X: in.A + in.B})
		}
	}
}
