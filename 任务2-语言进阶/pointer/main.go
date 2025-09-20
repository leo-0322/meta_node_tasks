package main

import "fmt"

func addTen(val *int) {
	*val += 10
}

func multipliedTwo(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

func main() {
	val := 2
	addTen(&val)
	fmt.Println(val)

	arr := []int{1, 2, 3}
	multipliedTwo(&arr)
	fmt.Println(arr)
}
