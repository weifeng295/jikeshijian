package main

import (
	"context"
	"math/rand"
	"time"
)

const (
	Success = iota
	Failure
	Timeout
	Rejection
	Size
)

// Event ...
type Event struct {
	Status int
}

// NewRandomEventStream ...
func NewRandomEventStream(ctx context.Context, duration time.Duration) <-chan *Event {
	events := make(chan *Event, 10)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(events)
				return
			default:
				events <- &Event{
					Status: rand.Intn(Size),
				}
			}
			time.Sleep(duration)
		}
	}()
	return events
}
