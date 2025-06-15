package main

// import "fmt"

// func getMessagesCostArrays() {

// 	messages, costs := getMessageWithRetries("hello", "world", "1")
// 	fmt.Println("Messages:", messages)
// 	fmt.Println("Costs:", costs)
// }

// func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
// 	arr := [3]string{primary, secondary, tertiary}
// 	var arr2 [3]int
// 	cost := 0

// 	for i, msg := range arr {
// 		cost += costcalculator(msg)
// 		arr2[i] = cost
// 	}

// 	return arr, arr2
// }

// func costcalculator(value string) int {
// 	return len(value)
// }
