package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Read() int {
	c.mu.RLock()             // Allow multiple readers
	defer c.mu.RUnlock()
	return c.value
}

func (c *Counter) Write(val int) {
	c.mu.Lock()              // Only one writer at a time
	defer c.mu.Unlock()
	c.value = val
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Start writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter.Write(42)
		fmt.Println("Written: 42")
	}()

	// Start 5 reader goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // Let the writer go first
			val := counter.Read()
			fmt.Printf("Reader %d read: %d\n", id, val)
		}(i)
	}

	wg.Wait()
}
