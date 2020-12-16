package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	services2 "go_study/gokit/gokit_grpc/services"
	"go_study/gokit/gokit_http/services"
	"strings"
)

func MakeArithmeticEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*services2.GrpcRequest)

		var (
			a, b, res int
			calError  error
		)

		a = int(req.A)
		b = int(req.B)

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Substract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, errors.New("the dividend can not be zero!")
		}
		if calError != nil {
			return &services2.GrpcResponse{Result: int32(res), Error: calError.Error()}, nil
		}
		return &services2.GrpcResponse{Result: int32(res), Error: ""}, nil
	}
}
