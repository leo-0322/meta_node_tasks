package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	// fmt.Println("producer start")
	for i := 0; i <= 100; i++ {
		ch <- i
		// fmt.Println("produced:", i)
		time.Sleep(10 * time.Millisecond)
	}
	// fmt.Println("producer end")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println("consumer start")
	for num := range ch {
		fmt.Println("consumed:", num)
		time.Sleep(15 * time.Millisecond)
	}
	// fmt.Println("consumer end")
}

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
	fmt.Println("main end")
}
