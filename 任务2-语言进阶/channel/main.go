package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	fmt.Println("producer start")
	for i := 0; i <= 10; i++ {
		ch <- i
		fmt.Println("produced:", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	fmt.Println("producer end")
}

func consumer(ch <-chan int) {
	fmt.Println("consumer start")
	for num := range ch {
		fmt.Println("consumed:", num)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println("consumer end")
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(3 * time.Second)
	fmt.Println("main end")
}
