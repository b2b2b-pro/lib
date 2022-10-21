package repo_srv

import (
	"context"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"

	"go.uber.org/zap"
)

func (gs *Controller) NewEntity(ctx context.Context, frm *torepo.Entity) (*torepo.NewEntityReply, error) {
	zap.S().Debugf("(s *GRPCServer)NewEntity. frm: %v\n", frm)

	var reply *torepo.NewEntityReply
	ent := object.MEntity(frm)

	r, err := gs.db.CreateEntity(*ent)
	if err != nil {
		reply = &torepo.NewEntityReply{Id: int32(r), Err: err.Error()}
	} else {
		reply = &torepo.NewEntityReply{Id: int32(r), Err: ""}
	}

	zap.S().Debugf("(s *GRPCServer)NewEntity. return: %v, %v\n", r, err)

	return reply, nil
}

func (gs *Controller) ListEntity(z *torepo.Zero, stream torepo.B2B2BService_ListEntityServer) error {
	zap.S().Debugf("(s *GRPCServer)ListEntity\n")

	var err error

	le, err := gs.db.ListEntity()
	if err != nil {
		zap.S().Errorf("gs.db.ListEntity() error: %v\n", err)
	}

	for _, x := range le {
		if err = stream.Send(object.GEntity(&x)); err != nil {
			return err
		}
	}

	return err
}
