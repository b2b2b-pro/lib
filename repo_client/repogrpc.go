/*
repo_client - gRPC-клиент к репозиторию
*/
package repo_client

import (
	"fmt"

	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RepoGRPC struct {
	client torepo.B2B2BServiceClient
	conn   *grpc.ClientConn
}

type ConfigRPC struct {
	HostRPC string
	PortRPC string
}

// TODO всё переделать, ибо просто проверка
func New(cfg *ConfigRPC) (*RepoGRPC, error) {
	var err error
	gc := &RepoGRPC{}
	url := fmt.Sprintf("%s:%s", cfg.HostRPC, cfg.PortRPC)

	gc.conn, err = grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Fatalf("did not connect: %v", err)
	}

	gc.client = torepo.NewB2B2BServiceClient(gc.conn)

	return gc, err
}

func (gc *RepoGRPC) Stop() {
	gc.conn.Close()
}
