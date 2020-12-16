package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go_study/gokit/gokit_log/services"
	"strings"
)

//接收http客户端的请求后，把请求参数转为请求模型对象，用于后续业务逻辑处理
type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

//用于向客户端响应结果
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

func MakeArithmeticEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		}

		return ArithmeticResponse{Result: res, Error: calError}, nil
	}
}
