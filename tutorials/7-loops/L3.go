package main

import "fmt"

func SendCostMultipliedMessages() {
	arr := []int{10, 35, 40, 55, 70}
	arr[7] = 4 
	for i := 0; i < len(arr); i++ {
		cost := 0.0
		cost += getMaxMessagesToSend(cost, arr[i])
		fmt.Printf("Total number of %d messages can be sent for $%d\n", thresh, arr[i])
	}
}
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
