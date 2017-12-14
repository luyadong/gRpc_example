package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRpc_example/cf"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
	ServerCert = "cert/server.crt"
	ServerKey  = "cert/server.key"
	CaCert = "cert/ca.crt"
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

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(ServerCert, ServerKey)
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CaCert)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	// Create the channel to listen on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not list on %s: %s", port, err)
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	// Create the gRPC server with the credentials
	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
