package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"go_study/gokit/gokit_grpc/endpoints"
	pb "go_study/gokit/gokit_grpc/services"
	"go_study/gokit/gokit_grpc/transports"
	"go_study/gokit/gokit_http/services"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc services.Service
	svc = services.ArithmeticService{}
	endpoint := endpoints.MakeArithmeticEndpoint(svc)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	endpoint = endpoints.MakeArithmeticEndpoint(svc)
	grpcServer := transports.MakeGRPCHandler(ctx, endpoint)

	go func() {
		//启动grpc服务
		ls, _ := net.Listen("tcp", ":9002")
		gs := grpc.NewServer()
		pb.RegisterGrpcServiceServer(gs, grpcServer)
		gs.Serve(ls)

	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
