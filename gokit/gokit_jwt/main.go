package main

import (
	"github.com/go-kit/kit/log"
	"go_study/gokit/gokit_jwt/endpoints"
	"go_study/gokit/gokit_jwt/services"
	"go_study/gokit/gokit_jwt/transports"
	"net/http"
	"os"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var svc services.Service
	svc = services.JwtService{}
	// add logging middleware
	svc = services.LogMiddlewareServer(logger)(svc)
	endpoints := endpoints.NewEndPointServer(svc, logger)
	httpHandler := transports.MakeHttpHandler(endpoints, logger)
	_ = http.ListenAndServe("0.0.0.0:9000", httpHandler)

}
