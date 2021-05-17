package main

import (
	"fmt"
	"runtime"
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
		buff := make([]byte, 1024)
		runtime.Stack(buff, false)
		fmt.Printf("panic with value: %v, %s\n", v, buff)
	}
}

func main() {
	callMakePanic()
	fmt.Println("after callMakePanic")
}
