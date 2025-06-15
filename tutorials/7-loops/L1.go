package main

import "fmt"

func SendBulkSMS() {
	arr := []int{10, 35, 40, 55, 70}

	for i := 0; i < len(arr); i++ {
		// numMessages := arr[i]
		// totalCost := bulkSend(numMessages)
		// fmt.Printf("Total cost for sending '%d' messages is $%.2f\n", numMessages,totalCost)

		thresh := withoutCondition(float64(arr[i]))
		fmt.Printf("Total number of %d messages can be sent for $%d\n", thresh,arr[i])
	} 
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}


func withoutCondition(threshold float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {

		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > threshold {
			return i
		}
	}
}