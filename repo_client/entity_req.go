package torepo_client

import (
	"context"
	"io"
	"time"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
)

func (gc *RepoGRPC) CreateEntity(frm object.Entity) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := gc.client.NewEntity(ctx, object.GEntity(&frm))
	if err != nil {
		zap.S().Fatalf("GRPC Client NewEntity error: %v", err)
	}

	zap.S().Debugf("GRPC New Entity Reply: %v", r.GetId())

	return int(r.GetId()), err
}

func (gc *RepoGRPC) ListEntity() ([]object.Entity, error) {
	r := []object.Entity{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := gc.client.ListEntity(ctx, &torepo.Zero{})
	if err != nil {
		zap.S().Fatalf("ListEntity error: %v\n", err)
	}

	for {
		x, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			zap.S().Fatalf("Loop ListEntity error: %v\n", err)
		}
		r = append(r, *object.MEntity(x))
		zap.S().Debugf("Loop stream recv: %v\n", x.Id)
	}

	return r, err
}
