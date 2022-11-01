package repo_client

import (
	"context"
	"io"
	"time"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
)

func (gc *RepoGRPC) CreateEntity(tkn string, frm object.Entity) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := gc.client.NewEntity(ctx, object.GEntity(tkn, &frm))
	if err != nil {
		zap.S().Debugf("GRPC Client NewEntity error: %v", err)
		return 0, err
	}

	zap.S().Debugf("GRPC New Entity Reply Id: %v\n", r.GetId())

	return int(r.GetId()), err
}

func (gc *RepoGRPC) ListEntity(tkn string) ([]object.Entity, error) {
	r := []object.Entity{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := gc.client.ListEntity(ctx, &torepo.Tkn{Tkn: tkn})
	if err != nil {
		zap.S().Debugf("ListEntity error: %v\n", err)
		return nil, err
	}

	for {
		x, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			zap.S().Debugf("Loop ListEntity error: %v\n", err)
		}

		r = append(r, *object.MEntity(x))
		zap.S().Debugf("Loop stream recv: %v\n", x.Id)
	}

	return r, err
}
