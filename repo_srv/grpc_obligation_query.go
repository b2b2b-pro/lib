package repo_srv

import (
	"context"

	"github.com/b2b2b-pro/lib/object"
	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
)

func (gs *Controller) NewObligation(ctx context.Context, obl *torepo.Obligation) (*torepo.NewObligationReply, error) {
	zap.S().Debugf("(s *GRPCServer)NewObligation. frm: %v\n", obl)

	var reply *torepo.NewObligationReply

	ent := object.MObligation(obl)

	r, err := gs.db.CreateObligation(obl.Tkn, *ent)
	if err != nil {
		reply = &torepo.NewObligationReply{Id: int32(r), Err: err.Error()}
	} else {
		reply = &torepo.NewObligationReply{Id: int32(r), Err: ""}
	}

	zap.S().Debugf("(s *GRPCServer)NewObligation. return: %v, %v\n", r, err)

	return reply, nil
}

func (gs *Controller) ListObligation(tkn *torepo.Tkn, stream torepo.B2B2BService_ListObligationServer) error {
	zap.S().Debugf("(s *GRPCServer)ListObligation\n")

	var err error

	le, err := gs.db.ListObligation(tkn.Tkn)
	if err != nil {
		zap.S().Errorf("gs.db.ListObligation() error: %v\n", err)
	}

	for _, x := range le {
		if err = stream.Send(object.GObligation(tkn.Tkn, &x)); err != nil {
			return err
		}
	}

	return err
}
