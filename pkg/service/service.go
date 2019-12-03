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
	GetDisputesByStatus(status string) (dispute.Disputes, error)
	ResolveDispute(dispute.Resolution) (bool, error)
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

func (s *service) GetDisputesByStatus(status string) (dispute.Disputes, error) {
	_, err := s.authorizor.IsModerator()
	if err != nil {
		return dispute.Disputes{}, err
	}
	return s.store.GetDisputesByStatus(status)
}

func (s *service) ResolveDispute(resultion dispute.Resolution) (bool, error) {
	_, err := s.authorizor.IsModerator()
	if err != nil {
		return false, err
	}
	return s.store.ResolveDispute(resultion)
}
