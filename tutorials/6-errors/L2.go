package main

import (
	"fmt"
)

func DividingError() {
	dividend := 10.0
	divisor := 0.0

	result, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result of dividing %.2f by %.2f is %.2f\n", dividend, divisor, result)
}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string { // implementing the default error interface because it has a method called Error same as the error interface
	return fmt.Sprintf("cannot divide %.2f by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
