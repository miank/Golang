package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second): // simulate work
		fmt.Println("work completed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	err := doWork(ctx)

	if err != nil {
		fmt.Println("work canceled:", err)
	} else {
		fmt.Println("work finished successfully")
	}
}
