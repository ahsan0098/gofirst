package main

import "fmt"

func main() {
	// defertest(true, true)
	// defertest(false, true)
	// defertest(true, false)
	// defertest(false, false)

	// test scope.go
	v1,v2 := splitEmail("drogon@dragonstone.com")

	fmt.Println(v1,v2)
}