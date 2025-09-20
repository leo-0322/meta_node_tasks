package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncMutex struct {
	mu    sync.Mutex
	count int
}

func (s *SyncMutex) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *SyncMutex) GetCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func main() {
	var wg sync.WaitGroup
	syncMutex := &SyncMutex{}

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				syncMutex.Increment()
			}
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	wg.Wait()
	dur := time.Since(startTime)
	fmt.Printf("All workers done in %v\n", dur)
	fmt.Printf("Final count: %d\n", syncMutex.GetCount())
}
