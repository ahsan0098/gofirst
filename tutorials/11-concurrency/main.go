package main

import (
	"time"
)

func main() {

	// send and receive emails orderly
	// l1();

	// channels practice
	// L2()

	// l3()

	// l4()

	// L5()
	// l6()

	// l7()

	// l8
	// logChan := make(chan string)
	// ticker := make(chan time.Time)
	// saveAfter := make(chan time.Time)

	// // Start the backup service
	// go saveBackups(ticker, saveAfter, logChan)

	// // Start a goroutine to simulate logs
	// go func() {
	// 	for msg := range logChan {
	// 		fmt.Println("[Log]", msg)
	// 	}
	// }()

	// // Simulate ticker snapshots
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		ticker <- time.Now()
	// 	}
	// }()

	// // Simulate save-after signal (e.g., after 5s)
	// time.Sleep(5 * time.Second)
	// saveAfter <- time.Now()

	// // Allow time to drain logs
	// time.Sleep(2 * time.Second)

	l9()
}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return // end the function after saving snapshot
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
