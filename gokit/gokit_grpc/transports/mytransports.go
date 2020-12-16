package transports

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"go_study/gokit/gokit_grpc/services"
)

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

type GrpcServer struct {
	grpcHandler grpc.Handler
}

func (g *GrpcServer) Add(ctx context.Context, in *services.GrpcRequest) (*services.GrpcResponse, error) {
	_, rsp, err := g.grpcHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*services.GrpcResponse), err

}
func (g *GrpcServer) Subtract(ctx context.Context, in *services.GrpcRequest) (*services.GrpcResponse, error) {
	_, rsp, err := g.grpcHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*services.GrpcResponse), err

}
func (g *GrpcServer) Multiply(ctx context.Context, in *services.GrpcRequest) (*services.GrpcResponse, error) {
	_, rsp, err := g.grpcHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*services.GrpcResponse), err

}
func (g *GrpcServer) Divide(ctx context.Context, in *services.GrpcRequest) (*services.GrpcResponse, error) {
	_, rsp, err := g.grpcHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*services.GrpcResponse), err

}

func MakeGRPCHandler(ctx context.Context, endpoint endpoint.Endpoint) *GrpcServer {
	return &GrpcServer{
		grpcHandler: grpc.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		),
	}
}
