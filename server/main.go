package main

import (
	"context"
	"github.com/mtojek/grpc-gateway-errors/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {}

func (hs *HelloServer) SayHello(ctx context.Context, req *api.HelloReq) (*api.HelloResp, error) {
	return nil, status.Error(codes.Unimplemented, "This is not implemented yet")
}

func main() {
	ctx := context.Background()

	httpPort := "50001"
	grpcPort := "50002"

	go func() {
		_ = RunServer(ctx, grpcPort, httpPort)
	}()

	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, new(HelloServer))

	conn, err := net.Listen("tcp", "localhost:" + grpcPort)
	if err != nil {
		log.Fatalf("Error occurred while starting listener: %v", err)
	}

	log.Fatal(s.Serve(conn))
}