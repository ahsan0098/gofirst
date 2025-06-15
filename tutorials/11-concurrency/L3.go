package main

import "fmt"

func l3() {
	numDBs := 3
	dbChan, count := getDBsChannel(numDBs)

	waitForDBs(numDBs, dbChan)

	fmt.Printf("All %d databases are online (count=%d)\n", numDBs, *count)
}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan // receive one signal per DB
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{} // signal DB is ready
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
