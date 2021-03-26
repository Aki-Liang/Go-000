package main

import (
	"flag"
	"log"
	"net"

	"github.com/Aki-Liang/Go-000/WeekFinal/finalproj/api"
	"github.com/Aki-Liang/Go-000/WeekFinal/finalproj/internal/server"
	"google.golang.org/grpc"
)

var (
	Version  string
	confPath string
	port     string
)

func init() {
	flag.StringVar(&confPath, "conf", "../../configs", "config path, eg: -conf=./config.json")
	flag.StringVar(&port, "port", ":80001", "server listen port, eg: -port=:80001")

}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return
	}

	stats := server.NewStatsServer()
	s := grpc.NewServer()
	api.RegisterStatsSvrService(s, stats)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
