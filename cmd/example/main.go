package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

const defaultLifetime = 1 * time.Second
const defaultEnvironment = "production"

type Config struct {
	Environment string
	Lifetime    time.Duration
}

func main() {
	c := new(Config)
	flag.StringVar(&c.Environment, "env", defaultEnvironment, "runtime environment type")
	flag.DurationVar(&c.Lifetime, "lt", defaultLifetime, "application lifetime")
	flag.Parse()

	t := time.NewTimer(time.Second)
	ctx, cncl := context.WithDeadline(
		context.Background(),
		time.Now().Add(c.Lifetime))
	defer cncl()

	for {
		select {
		case <-t.C:
			fmt.Println("Hello from", c.Environment)
			t.Reset(time.Second)
		case <-ctx.Done():
			fmt.Println("Goodbye")
			return
		}
	}
}
