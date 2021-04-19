package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func work(i, j int) error {
	fmt.Println("working", i, j)
	if rand.Int()%2 == 0 {
		return fmt.Errorf("bla bla bla error %d %d", i, j)
	}
	return nil
}

func main() {
	errgroup := errgroup.Group{}

	for i := 0; i < 5; i++ {
		var iCopy = i
		errgroup.Go(func() error {
			fmt.Printf("iCopy=%p i=%p\n", &iCopy, &i)
			return work(iCopy, i)
		})
	}

	if err := errgroup.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("main done")
}
