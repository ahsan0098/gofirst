package main

// import "fmt"

// func mailReport() {
// 	// Create an email with subscription
// 	email1 := email{
// 		isSubscribed: true,
// 		body:         "Hello, world.", // length 13, cost 2 13*2=26
// 	}

// 	// Create an email without subscription
// 	email2 := email{
// 		isSubscribed: false,
// 		body:         "Hello, world.", // length 13, cost 5 13*5=65
// 	}

// 	// Print reports
// 	printMailReport(email1,email1)
// 	printMailReport(email2,email2)
// }
// func printMailReport(e expense, f formatter) {
// 	fmt.Printf("Email cost: %d, Format: %s\n", e.cost(), f.format())
// }

// func (e email) cost() int {
// 	if e.isSubscribed {
// 		return 2 * len(e.body)
// 	}
// 	return 5 * len(e.body)
// }

// func (e email) format() string {
// 	if e.isSubscribed {
// 		return fmt.Sprintf("'%s' | Subscribed",e.body)
// 	}
// 	return fmt.Sprintf("'%s' | Not Subscribed",e.body)
// }

// type expense interface {
// 	cost() int
// }

// type formatter interface {
// 	format() string
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// }
