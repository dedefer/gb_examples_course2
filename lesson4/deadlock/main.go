package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println("goroutine 1")
		ch1 <- 1
	}()

	go func() {
		defer close(done)
		fmt.Println("goroutine 2")
		ch2 <- 1
	}()

	go func() {
		<-done
		fmt.Println("goroutine 0")
	}()

	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}
