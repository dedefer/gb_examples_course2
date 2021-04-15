package main

var arr = []int{}

func bar(i int) int {
	return arr[i]
}

func foo(x int) {
	bar(x + 1)
}

func main() {
	for i := 0; i < 10; i++ {
		foo(i)
	}
}
