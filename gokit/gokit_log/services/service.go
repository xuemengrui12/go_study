package services

type Service interface {
	Add(a, b int) int
}

type ArithmeticService struct {
}

func (s ArithmeticService) Add(a, b int) int {
	return a + b
}
