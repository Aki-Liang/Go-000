package rolling

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Bucket struct {
	Success  int64
	Failure  int64
	LastTime time.Time
}

func (b *Bucket) String() string {
	return fmt.Sprintf("[%s.%02d]bucket: success:%d failure:%d",
		b.LastTime.Format("2006-01-02 15:04:05"),
		b.LastTime.Nanosecond()/1e7,
		b.Success,
		b.Failure,
	)
}

type Window struct {
	buckets []*Bucket
	width   time.Duration
	cap     int
	idx     int
	wndLock sync.Mutex
}

func NewSlidingWindow(w time.Duration, cap int) *Window {
	buckets := []*Bucket{}
	now := time.Now()
	for i := 0; i < cap; i++ {
		buckets = append(buckets, &Bucket{
			LastTime: now,
			Success:  0,
			Failure:  0,
		})
	}

	return &Window{
		buckets: buckets,
		width:   w,
		cap:     cap,
		idx:     0,
	}
}

func (w *Window) getWnd() *Bucket {
	now := time.Now()
	w.wndLock.Lock()
	defer w.wndLock.Unlock()
	currBucket := w.buckets[w.idx]

	if now.Add(-w.width).After(currBucket.LastTime) {
		w.idx = (w.idx + 1) % w.cap
		currBucket = &Bucket{
			LastTime: now,
			Success:  0,
			Failure:  0,
		}
		w.buckets[w.idx] = currBucket
	}

	return currBucket
}

func (w *Window) AddSuccess() {
	b := w.getWnd()
	atomic.AddInt64(&b.Success, 1)
}

func (w *Window) AddFailure() {
	b := w.getWnd()
	atomic.AddInt64(&b.Failure, 1)
}

func (w *Window) PrintSum() {
	w.wndLock.Lock()
	defer w.wndLock.Unlock()

	var totalSuccess int64
	var totalFailure int64
	for _, b := range w.buckets {
		totalFailure += b.Failure
		totalSuccess += b.Success
	}

	fmt.Printf("total success:%v failure:%v", totalSuccess, totalFailure)
}

func (w *Window) PrintBuckets() {
	w.wndLock.Lock()
	defer w.wndLock.Unlock()
	for _, b := range w.buckets {
		fmt.Println(b)
	}
}
