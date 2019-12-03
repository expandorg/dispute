package service

import (
	"github.com/gemsorg/dispute/pkg/authentication"
	"github.com/gemsorg/dispute/pkg/authorization"
	"github.com/gemsorg/dispute/pkg/datastore"
	"github.com/gemsorg/dispute/pkg/dispute"
)

type DisputeService interface {
	Healthy() bool
	SetAuthData(data authentication.AuthData)
	CreateDispute(dispute.Dispute) (dispute.Dispute, error)
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

func (s *service) CreateDispute(d dispute.Dispute) (dispute.Dispute, error) {
	return s.store.CreateDispute(d)
}

func (s *service) SetAuthData(data authentication.AuthData) {
	s.authorizor.SetAuthData(data)
}
