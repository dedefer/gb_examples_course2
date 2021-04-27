package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := &User{Username: "Danila", Password: "123"}

	fmt.Printf("%+v\n", u)

	buff, err := proto.Marshal(u)
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", buff)

	u2 := &User{}

	proto.Unmarshal(buff, u2)

	fmt.Printf("%+v\n", u2)
}
