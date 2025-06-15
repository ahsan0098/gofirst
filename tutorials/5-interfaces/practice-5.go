package main

import "fmt"

func EmailOrMsgReportWithSwitch() {
	// Create an email with subscription
	email1 := email{
		isSubscribed: true,
		body:         "Hello, world.", // length 13, cost 0.13
		toAddress:    "sheikhahsanali98@gmail.com",
	}

	// Create an email without subscription
	email2 := email{
		isSubscribed: false,
		body:         "Hello, world.", // length 13, cost 0.65
		toAddress:    "sheikhahsanali98@gmail.com",
	}

	// Create an SMS with subscription
	sms1 := sms{
		isSubscribed:  true,
		body:          "Hello, world.", // length 13, cost 0.39
		toPhoneNumber: "1234567890",
	}

	// Create an SMS without subscription
	sms2 := sms{
		isSubscribed:  false,
		body:          "Hello, world.", // length 13, cost 1.3
		toPhoneNumber: "0987654321",
	}

	invalid := invalid{}
	// Print reports

	receiver, cost := getExpenseReport(email1)
	fmt.Printf("Email 1 to: %s , Cost: %.2f \n\n", receiver, cost)

	receiver, cost = getExpenseReport(email2)
	fmt.Printf("Email 2 to: %s , Cost: %.2f \n\n", receiver, cost)

	receiver, cost = getExpenseReport(sms1)
	fmt.Printf("SMS 1 to: %s , Cost: %.2f \n\n", receiver, cost)

	receiver, cost = getExpenseReport(sms2)
	fmt.Printf("SMS 2 to: %s , Cost: %.2f \n\n", receiver, cost)

	receiver, cost = getExpenseReport(invalid)
	fmt.Printf("Invalid message to: %s , Cost: %.2f \n\n", receiver, cost)
}
func getExpenseReport(e expense) (string, float64) {
	
	switch v := e.(type) {
	case email:
		return v.toAddress, e.cost()
	case sms:
		return v.toPhoneNumber, e.cost()
	default:
		return "Unknown", e.cost()
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
