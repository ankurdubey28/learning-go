package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)

	// wait for few seconds toa allow processing to complete
	time.Sleep(3 * time.Second)
}

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("cancelled:", ctx.Err())
	}
}
