package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

const defaultLifetime = 5 * time.Second

type Config struct {
	Environment string
	Lifetime    time.Duration
}

func main() {
	c := new(Config)
	flag.StringVar(&c.Environment, "env", "production", "runtime environment type")
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
			fmt.Println("new tick in ", c.Environment)
			t.Reset(time.Second)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
