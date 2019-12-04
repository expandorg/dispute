package dispute

import "time"

const (
	InPorgress = "inporgress"
	Accepted   = "accepted"
	Rejected   = "rejected"
)

type Dispute struct {
	ID                uint64    `json:"id" db:"id"`
	DisputeMessage    string    `json:"dispute_message" db:"dispute_message"`
	ResolutionMessage string    `json:"resolution_message" db:"resolution_message"`
	Status            string    `json:"status" db:"status"`
	Active            bool      `json:"active" db:"active"`
	WorkerID          uint64    `json:"worker_id" db:"worker_id"`
	ResponseID        uint64    `json:"response_id" db:"response_id"`
	ScoreID           uint64    `json:"score_id" db:"score_id"`
	TaskID            uint64    `json:"task_id" db:"task_id"`
	JobID             uint64    `json:"job_id" db:"job_id"`
	VerifierID        uint64    `json:"verifier_id" db:"verifier_id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type Disputes []Dispute

type Resolution struct {
	ResponseID uint64 `json:"response_id"`
	DisputeID  string `json:"dispute_id"`
	Status     string `json:"status"`
	Message    string `json:"resolution_message"`
}
