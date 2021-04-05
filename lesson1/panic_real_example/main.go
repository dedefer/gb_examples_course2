package main

import (
	"fmt"
)

func makePanic() {
	panic("panic in makePanic")
}

func callMakePanic() (err error) {
	defer panicHandlerWithHardLogic(func(e error) { err = e })

	defer fmt.Println("defer 1 makePanic")
	defer fmt.Println("defer 2 makePanic")

	makePanic()
	fmt.Println("after makePanic")

	return nil
}

func panicHandlerWithHardLogic(setter func(error)) {
	if v := recover(); v != nil {
		setter(fmt.Errorf("panic occured: %v", v))
	}
}

func main() {
	callMakePanic()
	fmt.Println("after callMakePanic")
}
