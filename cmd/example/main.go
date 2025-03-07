package main

import (
	"context"
	"fmt"
	"time"
)

const deadlineSeconds = 5

func main() {
	t := time.NewTimer(time.Second)
	ctx, cncl := context.WithDeadline(
		context.Background(),
		time.Now().Add(time.Second*deadlineSeconds))
	defer cncl()

	for {
		select {
		case <-t.C:
			fmt.Println("new tick")
			t.Reset(time.Second)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
