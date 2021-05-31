package main

import "fmt"

type WP interface {
	Do(func())
	DoBatch(...func())
	Close()
}

func printNum(i int) func() {
	return func() { fmt.Println("i'm printing", i) }
}

func main() {
	var wp WP = NewWorkerPool(100)
	defer wp.Close()

	tasks := []func(){}

	for i := 0; i < 20; i++ {
		tasks = append(tasks, printNum(i))
	}

	wp.DoBatch(tasks[:len(tasks)/2]...)
	fmt.Println("================")
	wp.DoBatch(tasks[len(tasks)/2:]...)
}
