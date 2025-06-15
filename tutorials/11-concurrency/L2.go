package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func L2() {
	emails := [3]email{{body: "hello wold", date: time.Now()}, {body: "hey there", date: time.Now()}, {body: "how are you", date: time.Now()}}

	arr := checkEmailAge(emails)

	fmt.Print(arr)
}
func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
