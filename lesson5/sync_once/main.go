package main

import (
	"fmt"
	"sync"
)

type Connector struct {
	connection string
	init       sync.Once
}

func (c *Connector) DoQuery(query string) {
	c.init.Do(func() {
		fmt.Println("establishing connection")
		c.connection = "connected"
	})

	fmt.Println("do query", query)
}

func main() {
	conn := Connector{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn.DoQuery(fmt.Sprint("i'm goroutine number", i))
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
