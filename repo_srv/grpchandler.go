/*
repo_srv обслуживает запросы, приходящие по gRPC, вызывает соответствующие методы
работы с хранилищем
*/
package repo_srv

import (
	"fmt"
	"net"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ConfigRPC struct {
	HostRPC string
	PortRPC string
}

type ObligationRepo interface {
	CreateOrigin(object.Origin) (int, error)
	CreateEntity(object.Entity) (int, error)
	ListEntity() ([]object.Entity, error)
	CreateObligation(object.Obligation) (int, error)
	ListObligation() ([]object.Obligation, error)
	Stop()
}

type Controller struct {
	torepo.UnimplementedB2B2BServiceServer
	s      *grpc.Server
	db     ObligationRepo
	cfg    ConfigRPC
	notify chan error
	// serverCtx     context.Context
	// serverStopCtx context.CancelFunc
}

func New(cfg ConfigRPC, rp ObligationRepo) (*Controller, error) {
	zap.S().Debugf("GRPCServer New config: %v\n", cfg)
	gs := &Controller{
		db:     rp,
		cfg:    cfg,
		notify: make(chan error, 1),
	}
	return gs, nil
}

func (gs *Controller) Start() {
	url := fmt.Sprintf("%s:%s", gs.cfg.HostRPC, gs.cfg.PortRPC)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		zap.S().Fatalf("failed to listen: %v", err)
	}
	gs.s = grpc.NewServer()
	torepo.RegisterB2B2BServiceServer(gs.s, gs)

	go func() {
		zap.S().Debugf("server listening at %v", lis.Addr())
		gs.notify <- gs.s.Serve(lis)
	}()
}

func (gs *Controller) Notify() <-chan error {
	return gs.notify
}

func (gs *Controller) Stop() {
	gs.s.GracefulStop()
}
