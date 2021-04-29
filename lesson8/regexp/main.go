package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println(
		regexp.MustCompile(`^a*$`).MatchString("aa"),
		regexp.MustCompile(`^\s+$`).MatchString("   "),
		regexp.MustCompile(`^\d+$`).MatchString("123"),
		regexp.MustCompile(`^\w+$`).MatchString("asdasd"),
		regexp.MustCompile(`^\w+-\d+$`).MatchString("asdasd-123123"),
	)

	re := regexp.MustCompile(`^(\w+)-(\d+)$`)
	fmt.Println(re.FindAllStringSubmatch("asdasd-123123", -1))
	// fmt.Println(b)
}
