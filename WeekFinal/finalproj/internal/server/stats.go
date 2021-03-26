package server

import (
	"context"

	"github.com/Aki-Liang/Go-000/WeekFinal/finalproj/api"
)

var _ api.UnstableStatsSvrService = new(StatsServer)

type StatsServer struct {
	api.StatsSvrService
}

func NewStatsServer() *StatsServer {
	return &StatsServer{}
}

func (s *StatsServer) GetRankList(ctx context.Context, req *api.RankListReq) (rsp *api.RankListRsp, err error) {

	return &api.RankListRsp{}, nil
}
