package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	p = flag.String("p", "123", "it is for example")
)

type Config struct {
	timeout  time.Duration
	attempts uint
}

func NewConfig(prefix string) *Config {
	c := &Config{}

	flag.DurationVar(&c.timeout, prefix+".timeout", time.Second, "timeout")
	flag.UintVar(&c.attempts, prefix+".attempts", 3, "attempts")

	return c
}

func main() {
	c1 := NewConfig("app1")
	c2 := NewConfig("app2")

	flag.Parse()
	fmt.Printf("%+v, %+v\n", c1, c2)
}
