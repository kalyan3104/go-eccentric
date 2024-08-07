package main

import (
	"context"
	"fmt"
	"time"
)

// NewDelayedCancelContext creates a context with a cancellation function that
// delays cancellation by a specified duration.
func NewDelayedCancelContext(parent context.Context, delay time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	go func() {
		// Wait for the specified delay before invoking cancel
		time.Sleep(delay)
		cancel()
	}()
	return ctx, cancel
}

func main() {
	// Create a parent context
	parentCtx := context.Background()

	// Create a context with a delayed cancellation
	delay := 2 * time.Second
	ctx, cancel := NewDelayedCancelContext(parentCtx, delay)
	defer cancel() // Ensure the cancel function is called to free resources

	// Start a goroutine that does some work
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Work canceled")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond) // Simulate work
			}
		}
	}()

	// Wait for a while to see the delayed cancellation
	time.Sleep(5 * time.Second)
	fmt.Println("Main function completed")
}
