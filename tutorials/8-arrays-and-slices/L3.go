package main

import (
	"fmt"
)

func L3() {
	messages := []string{"hello", "world", "1"}
	costs := getMessageCosts(messages)

	fmt.Println("Messages:", messages)
	fmt.Println("Costs:", costs)
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages)) // Preallocate

	for i, msg := range messages {
		costs[i] = float64(len(msg)) * 0.01
	}

	return costs
}
