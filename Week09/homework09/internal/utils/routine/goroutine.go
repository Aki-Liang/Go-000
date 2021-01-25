package routine

import (
	"context"
	"fmt"
)

func GO(ctx context.Context, proc func(ctx context.Context)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Goroutine panic， err:", err)
			}
		}()
		proc(ctx)
	}()
}
