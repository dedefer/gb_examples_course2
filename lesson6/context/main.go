package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx0, cancel0 := context.WithCancel(context.Background())
	ctx1, cancel1 := context.WithCancel(ctx0)
	ctx2 := context.WithValue(ctx1, "geek", "brains")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		<-ctx0.Done()
		fmt.Println("ctx0 done", ctx0.Value("geek"))
	}()

	go func() {
		defer wg.Done()
		<-ctx2.Done()
		fmt.Println("ctx2 done", ctx2.Value("geek"))
	}()

	cancel0()
	wg.Wait()

	cancel1()
}
