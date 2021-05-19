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

func getStackTrace() string {
	trace := make([]byte, 1024)
	runtime.Stack(trace, false)
	return string(trace)
}

func Avg(sequence []int) (avg int, err error) {

	defer func() {
		panicValue := recover()
		if panicValue != nil {
			fmt.Printf("PANIC: %v\n%s", panicValue, getStackTrace())
			err = NewError(fmt.Sprintf("%v", panicValue))
		}
	}()

	sum := 0
	for _, elem := range sequence {
		sum += elem
	}
	return sum / len(sequence), nil
}

func main() {
	for i, data := range testData {
		avg, err := Avg(data)
		if err != nil {
			fmt.Printf("test %d error: %s\n", i, err.Error())
		} else {
			fmt.Printf("test %d result: avg=%d\n", i, avg)
		}
	}
	fmt.Println("all tests done")
}
