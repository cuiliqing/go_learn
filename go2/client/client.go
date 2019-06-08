package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "../probuf"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	defer conn.Close()
	client := pb.NewMathServiceClient(conn)

	resp1, err := client.MathCalcuRpc(context.Background(), &pb.MathRequest{
		M1: 210.0,
		M2: 5.0,
	})
	log.Printf("Res:%f", resp1)
}
