package services

import "errors"

type Service interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) (int, error)
}

type ArithmeticService struct {
}

func (s ArithmeticService) Add(a, b int) int {
	return a + b
}

func (s ArithmeticService) Subtract(a, b int) int {
	return a - b
}

func (s ArithmeticService) Multiply(a, b int) int {
	return a * b
}

func (s ArithmeticService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero!")
	}
	return a / b, nil
}
