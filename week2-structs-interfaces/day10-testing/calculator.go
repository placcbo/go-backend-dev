package calculator

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a / b), nil
}
