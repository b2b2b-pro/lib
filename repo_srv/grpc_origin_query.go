package repo_srv

import (
	"context"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"

	"go.uber.org/zap"
)

func (gs *Controller) NewOrigin(ctx context.Context, og *torepo.Origin) (*torepo.NewOriginReply, error) {
	zap.S().Debugf("(s *GRPCServer)NewOrigin. frm: %v\n", og)

	var reply *torepo.NewOriginReply
	om := object.MOrigin(og)

	r, err := gs.db.CreateOrigin(*om)
	if err != nil {
		reply = &torepo.NewOriginReply{Id: int32(r), Err: err.Error()}
	} else {
		reply = &torepo.NewOriginReply{Id: int32(r), Err: ""}
	}

	zap.S().Debugf("(s *GRPCServer)NewOrigin. return: %v, %v\n", r, err)

	return reply, nil
}
