package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/emohankrishna/go-grpc/calculator/proto"
)

var address string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorServiceClient(conn)
	doCalculate(client)
	doCalculatePrimeDecomposition(client)
	doCalculateAverage(client)
}
