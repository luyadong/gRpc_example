package main

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "gRpc_example/cf"
)

const (
	address     = "localhost:50051"
	defaultNum = 20
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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