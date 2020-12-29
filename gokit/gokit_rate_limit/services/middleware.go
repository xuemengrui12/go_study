package services

import (
	"github.com/go-kit/kit/log"
)

const ContextReqUUid = "req_uuid"

type ServiceMiddleware func(Service) Service

type logMiddlewareServer struct {
	logger log.Logger
	next   Service
}

func (l logMiddlewareServer) Add(a, b int) int {
	l.logger.Log("function", "Add", "a", a, "b", b)
	return l.next.Add(a, b)
}

func LogMiddlewareServer(log log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return logMiddlewareServer{
			logger: log,
			next:   next,
		}
	}
}
