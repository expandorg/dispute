package service

import (
	"github.com/gemsorg/dispute/pkg/authorization"
	"github.com/gemsorg/dispute/pkg/datastore"
)

type DisputeService interface {
	Healthy() bool
}

type service struct {
	store      datastore.Storage
	authorizor authorization.Authorizer
}

func New(s datastore.Storage, a authorization.Authorizer) *service {
	return &service{
		store:      s,
		authorizor: a,
	}
}

func (s *service) Healthy() bool {
	return true
}
