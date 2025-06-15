package main

func L5() {
	ch := make(chan int)

	go sendReports(5, ch)

	count :=countReports(ch)

	println(count)
}
func countReports(numSentCh chan int) int {
	total := 0
	for {
		val, ok := <-numSentCh
		if !ok {
			break
		}
		total += val
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
