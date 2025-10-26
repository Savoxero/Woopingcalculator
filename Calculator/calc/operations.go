package calc

import "errors"

func Add(a, b float64) float64 {
	c := a + b
	return c
}

func Subtract(a, b float64) float64 {
	c := a - b
	return c
}

func Multiply(a, b float64) float64 {
	c := a * b
	return c
}
func Divide(a, b float64) (float64, error) {

	if b == 0 {

		return 0, errors.New("cannot divide by zero")

	}
	return a / b, nil

}
