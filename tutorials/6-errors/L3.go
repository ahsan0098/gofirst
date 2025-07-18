package main

import (
	"errors"
	"fmt"
)

func errorsInterface() {
	// Example usage of the divide function
	result, err := divide(10.0, 10.0)
	if err != nil {
		// Handle the error
		println("Error:", err.Error())
		return
	}
	fmt.Printf("Result: %.2f", result)
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}
