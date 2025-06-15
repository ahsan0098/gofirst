package main

import "fmt"

func L3() {
	names := []string{"billy", "bob", "joe", "billy"}

	maps := getNameCounts(names)

	for k, v := range maps {
		fmt.Printf("%q: %v\n", k, v)
	}
}


func getNameCounts(names []string) map[rune]map[string]int {
	counts :=make(map[rune]map[string]int)
	for _,name := range names{
		firstChar := rune(name[0])
		_,ok := counts[rune(firstChar)]

		if !ok{
			counts[firstChar]= make(map[string]int)
		}
		counts[firstChar][name]++
	}

	return counts;
}
