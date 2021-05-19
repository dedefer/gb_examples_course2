package main

import (
	"fmt"
	"runtime"
)

var testData = [][]int{
	{1, 2, 3, 4, 5, 5, 5},
	{-1, -2, -3, -4, -5, -5, -5},
	{}, // expected panic
	{1},
	{0},
}

func panicHandler() {
	panicValue := recover()
	if panicValue != nil {
		trace := make([]byte, 1024)
		runtime.Stack(trace, false)
		fmt.Printf("PANIC: %v\n%s", panicValue, trace)
	}
}

func Avg(sequence []int) int {
	defer panicHandler()

	sum := 0
	for _, elem := range sequence {
		sum += elem
	}
	return sum / len(sequence)
}

func main() {
	for i, data := range testData {
		avg := Avg(data)
		fmt.Printf("test %d result: avg=%d\n", i, avg)
	}
	fmt.Println("all tests done")
}
