package di

import (
	"github.com/Aki-Liang/Go-000/Week04/code/internal/dao"
	"github.com/Aki-Liang/Go-000/Week04/code/internal/server/grpc"
	"github.com/Aki-Liang/Go-000/Week04/code/internal/service"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, grpc.New, NewApp))
}
