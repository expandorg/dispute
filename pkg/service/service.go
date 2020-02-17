package service

import (
	"github.com/expandorg/dispute/pkg/authentication"
	"github.com/expandorg/dispute/pkg/authorization"
	"github.com/expandorg/dispute/pkg/datastore"
	"github.com/expandorg/dispute/pkg/dispute"
)

type DisputeService interface {
	Healthy() bool
	SetAuthData(data authentication.AuthData)
	CreateDispute(dispute.Dispute) (dispute.Dispute, error)
	GetDisputesByStatus(status string) (dispute.Disputes, error)
	GetDisputesByWorkerID(id uint64) (dispute.Disputes, error)
	ResolveDispute(dispute.Resolution) (dispute.Resolved, error)
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
	// isMod, err := s.authorizor.IsModerator()
	// fmt.Println("MOD", isMod)
	// if !isMod || err != nil {
	// 	return dispute.Disputes{}, err
	// }
	return s.store.GetDisputesByStatus(status)
}

func (s *service) GetDisputesByWorkerID(id uint64) (dispute.Disputes, error) {
	return s.store.GetDisputesByWorkerID(id)
}

func (s *service) ResolveDispute(resolution dispute.Resolution) (dispute.Resolved, error) {
	return s.store.ResolveDispute(resolution)
}
