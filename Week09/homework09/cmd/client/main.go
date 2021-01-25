package main

import (
	"context"
	"flag"
	"fmt"
	"homework09/internal/model"
	"homework09/internal/tcp"
	"homework09/internal/utils/routine"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", "0.0.0.0:5060", "tcp connect address")

func main() {
	flag.Parse()
	client := tcp.NewTCPClient(*addr)

	ctx, cancle := context.WithCancel(context.Background())
	closeFunc, recvChan, sendChan, err := client.Start(ctx)
	if err != nil {
		fmt.Println("start tcp client failed, err =", err.Error())
		return
	}
	routine.GO(ctx, func(c context.Context) {
		for {
			time.Sleep(time.Second)
			sendChan <- model.TCPMsg{
				Content: fmt.Sprintf("Client msg:%v", time.Now().Format("2006-01-02 15:04:05")),
			}
			select {
			case <-c.Done():
				return
			default:
				continue
			}

		}
	})
	routine.GO(ctx, func(c context.Context) {
		for {
			select {
			case msg := <-recvChan:
				fmt.Println("recv msg:", msg)
			case <-c.Done():
				return
			default:
				continue
			}

		}
	})
	fmt.Printf("tcp client started")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	log.Printf("get a signal %s", s.String())
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP:
		closeFunc()
		cancle()
	}
	fmt.Printf("tcp client exit")
}
