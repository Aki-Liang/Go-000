package grpc

import (
	"log"
	"net"

	pb "github.com/Aki-Liang/Go-000/Week04/code/api/killer/v1"
	"google.golang.org/grpc"
)

func New(svr pb.KillerServer) (s *grpc.Server, err error) {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s = grpc.NewServer()
	pb.RegisterKillerServer(s, svr)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return
}
