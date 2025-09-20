package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Task func()

func runTasks(tasks []Task) {
	for i, task := range tasks {
		wg.Add(1)
		go func(t Task, index int) {
			defer wg.Done()
			start := time.Now()
			t()
			dur := time.Since(start)
			fmt.Printf("任务 %d 执行时间: %v\n", index+1, dur)
		}(task, i)
	}
	wg.Wait()
	fmt.Println("所有任务完成")
}

func printOddNum() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("打印奇数：%d\n", i)
		}
	}
}

func printEvenNum() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("打印偶数：%d\n", i)
		}
	}
}

func main() {
	// wg.Add(2)
	// go printEvenNum()
	// go printOddNum()
	// wg.Wait()
	// fmt.Println("main 结束")

	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务 1 执行完成")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务 2 执行完成")
		},
	}
	runTasks(tasks)
}
