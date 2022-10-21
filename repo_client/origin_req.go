package repo_client

import (
	"context"
	"time"

	"github.com/b2b2b-pro/lib/object"
	"go.uber.org/zap"
)

func (gc *RepoGRPC) CreateOrigin(om object.Origin) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := gc.client.NewOrigin(ctx, object.GOrigin(&om))
	if err != nil {
		zap.S().Debugf("GRPC Client NewOrigin error: %v", err)

		return 0, err
	}

	zap.S().Debugf("GRPC New Origin Reply: %v", r.GetId())

	return int(r.GetId()), err
}
