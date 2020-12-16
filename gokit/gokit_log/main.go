package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"go_study/gokit/gokit_log/endpoints"
	"go_study/gokit/gokit_log/services"
	"go_study/gokit/gokit_log/transports"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc services.Service
	svc = services.ArithmeticService{}
	// add logging middleware
	svc = services.LogMiddlewareServer(logger)(svc)
	endpoint := endpoints.MakeArithmeticEndpoint(svc)
	// add logging middleware
	endpoint = endpoints.LoggingMiddleware(logger)(endpoint)
	handler := transports.MakeHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
