package concurrency

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeoutsDeadlines(t *testing.T) {
	// Create a context with a timeout of 2 seconds !!!
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure the context is canceled when done

	// Simulate a long-running task
	go longRunningTask(ctx)

	// Wait for the task to complete or the timeout to expire
	select {
	case <-ctx.Done():
		// The context is done, either due to task completion or timeout
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Task took too long and timed out!")
		} else {
			fmt.Println("Task completed within the timeout.")
		}
	}
}

func longRunningTask(ctx context.Context) {
	// Simulate a task that takes some time
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed.")
	case <-ctx.Done():
		// The context was canceled or timed out
		fmt.Println("Task canceled or timed out.")
	}
}
