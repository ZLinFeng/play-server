package initial

import (
	"context"
	"fmt"

	"github.com/ZLinFeng/play-server/config"
	"github.com/ZLinFeng/play-server/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServerRpc(c *config.TransportConfig) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	host := "10.10.188.116:12314"
	conn, err := grpc.NewClient(host, opts...)
	if err != nil {
		fmt.Printf("grpc.NewClient err: %v\n", err)
	}
	defer conn.Close()
	client := rpc.NewHelloClient(conn)
	response, err := client.SayHello(context.Background(), &rpc.HelloRequest{Name: "Edison"})
	if err != nil {
		fmt.Printf("client.SayHello err: %v\n", err)
	}
	fmt.Printf("Response: %s\n", response.Message)
}
