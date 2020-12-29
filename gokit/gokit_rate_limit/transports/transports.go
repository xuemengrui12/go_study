package transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"go_study/gokit/gokit_rate_limit/endpoints"
	"go_study/gokit/gokit_rate_limit/services"
	"net/http"
	"strconv"
)

//把用户的请求内容转换为请求对象
func decodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, err := strconv.Atoi(pa)
	if err != nil {
		return nil, ErrorBadRequest
	}
	b, err := strconv.Atoi(pb)
	if err != nil {
		return nil, ErrorBadRequest
	}
	return endpoints.ArithmeticRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}

//把处理结果转换为响应对象
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		//kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
		kithttp.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			logger.Log(fmt.Sprint(ctx.Value(services.ContextReqUUid)), err)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(err)
		}),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
			logger.Log("给请求添加uuid", UUID)
			ctx = context.WithValue(ctx, services.ContextReqUUid, UUID)
			return ctx
		}),
	}
	r.Methods("GET").Path("/calculate/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint,
		decodeArithmeticRequest,
		encodeArithmeticResponse,
		options...,
	))

	return r
}
