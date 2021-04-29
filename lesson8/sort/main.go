package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 1, 0}
	b := []string{"c", "b", "a", ""}

	sort.Slice(b, func(i, j int) bool {
		fmt.Println(i, j)
		return a[i] < a[j]
	})

	fmt.Println(a)
	fmt.Println(b)
}
