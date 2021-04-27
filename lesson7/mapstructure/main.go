package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Username string `mapstructure:"user"`
	Password string
}

func main() {
	user := map[interface{}]interface{}{
		"user":     "danila",
		"password": "fomin",
	}

	u := &User{}

	mapstructure.Decode(user, u)

	user1 := map[string]string{}

	mapstructure.Decode(u, &user1)

	fmt.Printf("%+v\n", user1)
}
