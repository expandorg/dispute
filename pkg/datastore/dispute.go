package datastore

import (
	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	CreateDispute(dispute.Dispute) (dispute.Dispute, error)
	GetDisputesByStatus(status string) (dispute.Disputes, error)
	GetDisputesByWorkerID(id uint64) (dispute.Disputes, error)
	ResolveDispute(dispute.Resolution) (dispute.Resolved, error)
}

type DisputeStore struct {
	DB *sqlx.DB
}

func NewDisputeStore(db *sqlx.DB) *DisputeStore {
	return &DisputeStore{
		DB: db,
	}
}

func (s *DisputeStore) CreateDispute(d dispute.Dispute) (dispute.Dispute, error) {
	var newDisp dispute.Dispute
	result, err := s.DB.Exec(
		"INSERT INTO disputes (dispute_message, status, active, worker_id, response_id, score_id, task_id, job_id, verifier_id) VALUES (?,?,?,?,?,?,?,?,?)",
		d.DisputeMessage, dispute.Pending, true, d.WorkerID, d.ResponseID, d.ScoreID, d.TaskID, d.JobID, d.VerifierID)

	if err != nil {
		return newDisp, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return newDisp, err
	}

	disp, err := s.GetDispute(uint64(id))

	if err != nil {
		return newDisp, err
	}

	return disp, nil
}

func (s *DisputeStore) GetDispute(id uint64) (dispute.Dispute, error) {
	disp := dispute.Dispute{}

	err := s.DB.Get(&disp, "SELECT * FROM disputes WHERE id=?", id)

	if err != nil {
		return disp, err
	}
	return disp, nil
}

func (s *DisputeStore) GetDisputesByStatus(status string) (dispute.Disputes, error) {
	disp := dispute.Disputes{}
	query := "SELECT * FROM disputes WHERE active IS TRUE"
	var err error
	// if status is empty, then give me all statuses
	if status != "" {
		query = query + " AND status=?"
		err = s.DB.Select(&disp, query, status)
	} else {
		err = s.DB.Select(&disp, query)
	}

	if err != nil {
		return disp, err
	}
	return disp, nil
}

func (s *DisputeStore) GetDisputesByWorkerID(id uint64) (dispute.Disputes, error) {
	disp := dispute.Disputes{}

	err := s.DB.Select(&disp, "SELECT * FROM disputes WHERE worker_id=?", id)
	if err != nil {
		return disp, err
	}

	return disp, nil
}

func (s *DisputeStore) ResolveDispute(resolution dispute.Resolution) (dispute.Resolved, error) {
	var updated dispute.Resolved
	result, err := s.DB.Exec(
		"UPDATE disputes SET status=?, resolution_message=?, active=0 WHERE id=?",
		resolution.Status, resolution.Message, resolution.DisputeID)

	if err != nil {
		return updated, err
	}

	num, err := result.RowsAffected()
	if err != nil || num == 0 {
		return updated, err
	} else {
		updated.Success = true
	}

	return updated, nil
}
