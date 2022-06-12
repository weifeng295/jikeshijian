package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewSlidingWindow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	window := NewSlidingWindow(ctx, 1, 10)
	window.Start(NewRandomEventStream(ctx, time.Millisecond*1000))
	for summary := range window.Subscribe() {
		fmt.Printf("success: %d, failure %d, timeout: %d, rejection: %d, total: %d, success rate: %f \n",
			summary.Success,
			summary.Failure,
			summary.Timeout,
			summary.Rejection,
			summary.Total,
			float64(summary.Success)/float64(summary.Total),
		)
	}
}
