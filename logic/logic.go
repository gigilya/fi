package logic

import (
	"math"
)

type IFibonacci interface {
	IsFibonacci(num int) bool
	NearestFibonacci(num int) int
	AdjacentFibonacci(num int) (int, int)
}

type FibonacciService struct {
}

//проверяет, является ли данное число числом Фибоначчи
func (fs *FibonacciService) IsFibonacci(num int) bool {
	return fs.isPerfectSquare(5 * num * num + 4) || fs.isPerfectSquare(5*num*num-4)
}

//возвращает два числа Фибоначчи, ближайшие к данному числу
func (fs *FibonacciService) AdjacentFibonacci(num int) (int, int) {
	if num == 0 {
		return 0, 1
	}
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	return a, b + a
}

//возвращает ближайшее число Фибоначчи к данному числу
func (fs *FibonacciService) NearestFibonacci(num int) int {
	for i := 1; ; i++ {
		if fs.IsFibonacci(num + i) {
			return num + i
		}
		if fs.IsFibonacci(num - i) {
			return num - i
		}
	}
}

//проверяет, является ли данное число совершенным квадратом
func (fs *FibonacciService) isPerfectSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}