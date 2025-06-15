package main

import (
	"errors"
	"fmt"
)


func L2() {
	arr := [3]string{"pro","free","enterprise"}
	messages := [3]string{"hello", "world", "1"}

	for _,plan := range arr {
		slice, err := getMessageWithRetriesForPlan(plan, messages)
		if err != nil {
			fmt.Printf("error %v \n",err)
			continue
		}
		fmt.Printf("message %v \n",slice)
	}
}

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		return messages[:], nil
	}

	if plan == "free" {
		return messages[0:2], nil
	}

	return nil, errors.New("unsupported plan")

}
