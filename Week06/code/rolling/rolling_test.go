package rolling

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWindow(t *testing.T) {
	wnd := NewSlidingWindow(time.Second, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				n := rand.Intn(2)
				if n%2 == 0 {
					wnd.AddFailure()
				} else {
					wnd.AddSuccess()
				}
				wnd.PrintBuckets()
				time.Sleep(time.Duration(n) * time.Second)
			}

		}()
	}
	wg.Wait()
	wnd.PrintSum()
}
