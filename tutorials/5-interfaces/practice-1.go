package main

import (
	"fmt"
	"time"
)


func sendingBirthdayReportMs() {	// Create a birthday message
	birthdayMsg := birthdayMessage{
		birthdayTime:  time.Date(2023, time.October, 10, 0, 0, 0, 0, time.UTC),
		recipientName: "Alice",
	}

	// Create a sending report
	sendingReportMsg := sendingReport{
		reportName:    "Weekly Summary",
		numberOfSends: 42,
	}

	// Send messages
	sendMessage(birthdayMsg)
	sendMessage(sendingReportMsg)
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}