package di

import (
	"context"
	"time"

	"log"

	"github.com/Aki-Liang/Go-000/Week04/code/internal/service"
	"google.golang.org/grpc"
)

type App struct {
	svc     *service.Service
	grpcSvr *grpc.Server
}

func NewApp(svc *service.Service, grpcSvr *grpc.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:     svc,
		grpcSvr: grpcSvr,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := grpcSvr.Shutdown(ctx); err != nil {
			log.Printf("grpcSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
