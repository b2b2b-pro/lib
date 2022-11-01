package repo_client

import "github.com/b2b2b-pro/lib/object"

func (gc *RepoGRPC) CreateObligation(tkn string, obl object.Obligation) (int, error) {
	return 1, nil
}

func (gc *RepoGRPC) ListObligation(tkn string) ([]object.Obligation, error) {
	r := []object.Obligation{}
	return r, nil
}
