package main

import (
	"context"
	"log"
	"time"

	pb "newnok6/logic-test/pie-fire-dire/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMeatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.MeatRequest{}
	res, err := client.GetMeat(ctx, req)
	if err != nil {
		log.Fatalf("could not call method: %v", err)
	}

	log.Printf("Response: %v", res.GetBeef())
}
