package main

import (
	pb "../probuf"
	"context"
	"log"
	"net"

	 "google.golang.org/grpc"
)

type MathServiceServer struct {

}

func (m *MathServiceServer) MathCalcuRpc(ctx context.Context, in *pb.MathRequest) (res *pb.MathResponse, err error) {
	log.Printf("client data sent:%f-%f", in.M1, in.M2)
	if err != nil {
		return
	}
	var r = in.M1 /in.M2
	return &pb.MathResponse{
		Res:r,
	}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMathServiceServer(grpcServer, &MathServiceServer{})
	grpcServer.Serve(lis)

}

