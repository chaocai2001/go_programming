package concurrency

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelledByCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelByCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelledByCtx(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
