package services

import (
	"context"
	"errors"
	"go_study/gokit/gokit_jwt/utils"
)

type Service interface {
	Add(a, b int) int
	Login(ctx context.Context, request LoginRequest) (response LoginResponse, err error)
}

type JwtService struct {
}

func (s JwtService) Add(a, b int) int {
	return a + b
}

func (s JwtService) Login(ctx context.Context, request LoginRequest) (response LoginResponse, err error) {
	if request.Username != "xmr" || request.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	response.Token, err = utils.CreateJwtToken(request.Username, 1)
	return response, err
}
