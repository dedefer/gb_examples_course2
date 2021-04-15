// Docs module provides FooBar function

// 	FooBar(1)
package docs

import "fmt"

// FooBar is actually FizzBuzz
func FooBar(num int) {
	for i := 1; i <= num; i++ {
		str := ""
		if i%3 == 0 {
			str += "foo"
		}
		if i%5 == 0 {
			str += "bar"
		}
		if str == "" {
			str = fmt.Sprint(i)
		}

		fmt.Println(str)
	}
}
