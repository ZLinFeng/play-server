package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ZLinFeng/play-server/rpc"
	"google.golang.org/grpc"
)

type AppClient struct {
	rpc.UnimplementedHelloServer
}

func (s *AppClient) SayHello(c context.Context, req *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	res := fmt.Sprintf("Hello, %s", req.Name)

	return &rpc.HelloResponse{Message: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 12314))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)
	rpc.RegisterHelloServer(server, &AppClient{})
	server.Serve(lis)
}
