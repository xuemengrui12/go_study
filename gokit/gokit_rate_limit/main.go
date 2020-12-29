package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	//"go.uber.org/ratelimit"
	"go_study/gokit/gokit_rate_limit/endpoints"
	"go_study/gokit/gokit_rate_limit/services"
	"go_study/gokit/gokit_rate_limit/transports"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	svc = services.LogMiddlewareServer(logger)(svc)
	endpoint := endpoints.MakeArithmeticEndpoint(svc)
	// 使用x/time/rate
	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 3)
	endpoint = endpoints.NewGolangRateAllowMiddleware(ratebucket)(endpoint)
	//使用go.uber.org/ratelimit
	//uberLimit := ratelimit.New(1)         //一秒请求一次
	//endpoint = endpoints.NewUberRateMiddleware(uberLimit)(endpoint)
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
