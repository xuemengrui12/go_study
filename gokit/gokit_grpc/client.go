package main

import (
	"context"
	"fmt"
	pb "go_study/gokit/gokit_grpc/services"
	"google.golang.org/grpc"
)

func main() {
	serviceAddress := "127.0.0.1:9002"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	grpcClient := pb.NewGrpcServiceClient(conn)
	response, err := grpcClient.Add(context.Background(), &pb.GrpcRequest{A: 5, B: 4, RequestType: "Add"})
	fmt.Println(response.Result)
}
