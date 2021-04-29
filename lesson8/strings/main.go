package main

import (
	"fmt"
	"strings"
)

func main() {
	strings.Contains("abcd", "bc")   // true
	strings.HasPrefix("abcd", "abc") // true
	strings.HasSuffix("", "")

	fmt.Println(strings.Split("12345", ""))

	strings.ReplaceAll("aabaa", "a", "c")
	r := strings.NewReplacer(
		"a", "1",
		"b", "2",
		"c", "3",
	)

	fmt.Println(r.Replace("abcd"))
}
