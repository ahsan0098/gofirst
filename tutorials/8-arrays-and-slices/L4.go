package main

import "fmt"

func L4() {
	nums := []int{5, 10, 15, 20, 25}

	sums := sum(nums...);
	fmt.Println(sums)
}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
