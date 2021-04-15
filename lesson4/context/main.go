package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// go func() {
	// 	ticker := time.NewTicker(time.Second)
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			fmt.Println("new tick")
	// 		case <-ctx.Done():
	// 			fmt.Println("Done")
	// 			return
	// 		}
	// 	}
	// }()

	go func() {
		ticker := 0
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				ticker++
			}
		}
	}()

	var str string
	fmt.Scanln(&str)
	cancel()
	time.Sleep(time.Second)
}
