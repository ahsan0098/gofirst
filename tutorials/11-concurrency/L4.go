package main

import "fmt"

func l4() {
	emails := []string{"admin@gmail.com", "alex@gmail.com"}

	chann := addEmailsToQueue(emails)

	for i := 0; i < len(emails); i++ {
		mail := <-chann
		fmt.Printf("sending now %s \n",mail)
	}

	close(chann)
}
func addEmailsToQueue(emails []string) chan string {
	chanel := make(chan string, len(emails))
	// go func() {
		for _, mail := range emails {
			chanel <- mail
		}
	// }()

	return chanel
}
