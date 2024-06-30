package main

import (
	"context"
	"fmt"
	"time"
)

// Worker fetches and adds works to a remote work orchestration server.
type Worker struct{}

type Work struct{}

func (w *Worker) Fetch(ctx context.Context) (*Work, error) {
	_ = ctx
	ctx.Done()
	return &Work{}, nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
