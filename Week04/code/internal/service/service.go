package service

import (
	"context"
	"fmt"

	pb "homework04/api/killer/v1"
	"homework04/internal/dao"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.KillerServer), new(*Service)))

type Service struct {
	dao dao.Dao
}

func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

func (s *Service) Kill(ctx context.Context, req *pb.KillReq) (*pb.KillRsp, error) {
	rsp := &pb.KillRsp{}
	target, err := s.dao.Find(ctx, req.GetTargetName())
	if err != nil {
		return rsp, err
	}
	target.Status = "killd"
	rsp.Message = fmt.Sprintf("%v status->%v", target.TargetName, target.Status)
	return rsp, err
}

func (s *Service) Close() {
	return
}
