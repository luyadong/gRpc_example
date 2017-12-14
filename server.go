package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRpc_example/cf"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
	ServerCert = "cert/server.crt"
	ServerKey  = "cert/server.key"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.CfRequest) (*pb.CfReply, error) {
	rs := in.Num1 + in.Num2
	return &pb.CfReply{Sum: rs}, nil
}

func (s *server) Del(ctx context.Context, in *pb.CfRequest) (*pb.CfReply, error) {
	rs := in.Num1 - in.Num2
	return &pb.CfReply{Sum: rs}, nil
}


func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(ServerCert, ServerKey)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
