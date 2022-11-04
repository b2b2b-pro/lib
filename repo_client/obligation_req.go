package repo_client

import (
	"context"
	"io"
	"time"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
)

func (gc *RepoGRPC) CreateObligation(tkn string, obl object.Obligation) (int, error) {
	return 1, nil
}

func (gc *RepoGRPC) ListObligation(tkn string) ([]object.Obligation, error) {
	r := []object.Obligation{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := gc.client.ListObligation(ctx, &torepo.Tkn{Tkn: tkn})
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
			zap.S().Debugf("Loop ListObligation error: %v\n", err)
			break
		}

		r = append(r, *object.MObligation(x))
		zap.S().Debugf("Loop stream recv: %v\n", x.Id)
	}

	return r, err
}
