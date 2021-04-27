package main

import (
	"fmt"
)

func main() {
	u := &User{Username: "Danila", Password: "123"}

	fmt.Printf("%+v\n", u)

	buff, err := u.Marshal()
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", buff)

	u2 := &User{}

	u2.Unmarshal(buff)

	fmt.Printf("%+v\n", u2)
}
