package main

import (
	"fmt"
)

func SMSToCouple() {

	
	msgToCustomer := "Happy Anniversary!"
	msgToSpouse := "Happy Anniversary to you too!"

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total cost for sending messages: %d cents\n", totalCost)

}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	if len(msgToCustomer) == 0 || len(msgToSpouse) == 0 {
		return 0, fmt.Errorf("both messages must be non-empty")
	}

	costCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, fmt.Errorf("error sending message to customer: %w", err)
	}

	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, fmt.Errorf("error sending message to spouse: %w", err)
	}

	totalCost := costCustomer + costSpouse
	return totalCost, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
