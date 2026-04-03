package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

func process(ctx context.Context) {
	id := ctx.Value(requestIDKey).(string)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("request", id, "cancelled:", ctx.Err())
			return
		default:
			fmt.Println("processing request", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-1")

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go process(ctx)

	time.Sleep(3 * time.Second)
}
