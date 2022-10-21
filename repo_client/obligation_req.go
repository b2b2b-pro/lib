package torepo_client

import "github.com/b2b2b-pro/lib/object"

func (gc *RepoGRPC) CreateObligation(object.Obligation) (int, error) {
	return 1, nil
}

func (gc *RepoGRPC) ListObligation() ([]object.Obligation, error) {
	r := []object.Obligation{}
	return r, nil
}
