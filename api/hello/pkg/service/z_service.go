package service

import (
	"context"
)

// HelloService describes the service.
type HelloService interface {
	Hello(ctx context.Context, s string) (rs string, err error)
}

type basicHelloService struct{}

func (b *basicHelloService) Hello(ctx context.Context, s string) (rs string, err error) {
	return s, err
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
