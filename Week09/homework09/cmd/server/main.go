package main

import (
	"context"
	"flag"
	"fmt"
	"homework09/internal/tcp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var addr = flag.String("addr", "0.0.0.0:5060", "tcp listen address")

func main() {
	flag.Parse()
	svr := tcp.NewTCPServer(*addr)
	ctx, cancle := context.WithCancel(context.Background())
	err := svr.Start(ctx)
	if err != nil {
		fmt.Println("start tcp server failed, err =", err.Error())
		return
	}
	fmt.Printf("tcp server started")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP:
			cancle()
			fmt.Printf("tcp server exit")
			return
		}
	}
}
