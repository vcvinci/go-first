package cancel_by_close

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCannelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

/**
 * root  Context: context.Background()
 * child Context: context.WithCannel(parentContext)
 *		ctx, cancel := context.WithCancel(context.Backgroud())
 * current Context is canceled, it`s children context will be all canceled
 * receive cancel notify: <- ctx.Done()
 */
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCannelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	// 只会关闭一个
	//cancel_1(cancelChan)
	cancel()
	time.Sleep(time.Second * 1)
}
