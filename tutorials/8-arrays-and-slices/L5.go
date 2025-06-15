package main

import "fmt"

func L5() {
	day1 := cost{
		day:   1,
		value: 0.05,
	}
	day2 := cost{
		day:   2,
		value: 0.15,
	}
	day3 := cost{
		day:   1,
		value: 0.25,
	}

	costs := []cost{day1, day2, day3}

	for i := 1; i < 5; i++ {
		dayCosts := getDayCosts(costs, i)
		fmt.Printf("Costs for day %d: %v\n", i, dayCosts)
	}
}

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	var dayCosts []float64
	for _, c := range costs {
		if c.day == day {
			dayCosts = append(dayCosts, c.value)
		}
	}
	return dayCosts
}
