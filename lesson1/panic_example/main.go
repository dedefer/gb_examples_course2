package main

import (
	"fmt"
)

func makePanic() {
	panic("panic in makePanic")
}

func callMakePanic() {
	defer panicHandler()
	defer fmt.Println("defer 1 makePanic")
	defer fmt.Println("defer 2 makePanic")

	makePanic()
	fmt.Println("after makePanic")
}

func panicHandler() {
	if v := recover(); v != nil {
		fmt.Printf("panic with value: %v\n", v)
	}
}

func main() {
	callMakePanic()
	fmt.Println("after callMakePanic")
}
