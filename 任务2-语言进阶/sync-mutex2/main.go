package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var count int32

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait()
	dur := time.Since(startTime)
	fmt.Printf("All workers done in %v\n", dur)
	fmt.Printf("Final count: %d\n", count)
}
