package main

import "fmt"

func main() {

	messages := []Message{
		{Recipient: "alex@gmail.com", Success: true},
		{Recipient: "john@gmail.com", Success: false},
		{Recipient: "peter@gmail.com", Success: false},
		{Recipient: "tom@gmail.com", Success: true},
	}

	var analytics Analytics

	for _, message := range messages {
		getMessageText(&analytics, message)
	}

	fmt.Printf("you send %d messages among them %d were successful and %d were failed", analytics.MessagesTotal, analytics.MessagesSucceeded, analytics.MessagesFailed)
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(analyt *Analytics, m Message) {
	an := *analyt

	if !m.Success {
		an.MessagesFailed++
	}
	if m.Success {
		an.MessagesSucceeded++
	}

	an.MessagesTotal++

	*analyt = an
}
