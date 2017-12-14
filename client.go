package main

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRpc_example/cf"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultNum = 20
	ClientCert = "cert/client.crt"
	ClientKey  = "cert/client.key"
	CaCert = "cert/ca.crt"
)

func main() {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(ClientCert, ClientKey)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CaCert)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	// Create the TLS credentials for transport
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "server.maoyan.com", //server.crt文件签名的serverName
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect to %s: %s", address, err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	var num1,num2 int64
	num1 = defaultNum
	num2 = defaultNum

	if len(os.Args) > 1 {
		a,_ := strconv.ParseInt(os.Args[1],0,64)
		b,_ := strconv.ParseInt(os.Args[2],0,64)
		num1 = a
		num2 = b

	}

	r, err := c.Add(context.Background(), &pb.CfRequest{Num1: num1, Num2:num2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.Sum)
}