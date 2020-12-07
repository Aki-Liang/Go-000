package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	g, cancelCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return hookSignal(cancelCtx)
	})
	g.Go(func() error {
		return startServer(cancelCtx, ":8086", &httpHandler{})
	})
	if err := g.Wait(); err != nil {
		fmt.Println("error group return err:", err.Error())
	}
	filepath.Walk()

	fmt.Println("DONE!")
}

func hookSignal(ctx context.Context) error {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("signal routine：START!")
	for {
		select {
		case s := <-c:
			//发信号中止http server
			return fmt.Errorf("get %v signal", s)
		case <-ctx.Done():
			return fmt.Errorf("signal routine：other work done")
		}
	}
}

func startServer(ctx context.Context, addr string, h http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("http routione：other work done")
		s.Shutdown(context.Background())
	}(ctx)
	fmt.Println("http routione：START!")
	return s.ListenAndServe()
}

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {

}
