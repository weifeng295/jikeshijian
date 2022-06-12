package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewRandomEventStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	stream := NewRandomEventStream(ctx, time.Second)
	for e := range stream {
		fmt.Println(e)
	}
}
