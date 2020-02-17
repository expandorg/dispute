package server

import (
	"net/http"

	"github.com/expandorg/dispute/pkg/authentication"

	"github.com/jmoiron/sqlx"

	"github.com/expandorg/dispute/pkg/api/disputecreator"
	"github.com/expandorg/dispute/pkg/api/disputeresolver"
	"github.com/expandorg/dispute/pkg/api/disputesfetcher"
	"github.com/expandorg/dispute/pkg/api/healthchecker"
	"github.com/expandorg/dispute/pkg/service"
	"github.com/gorilla/mux"
)

func New(
	db *sqlx.DB,
	s service.DisputeService,
) http.Handler {
	r := mux.NewRouter()

	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/disputes", disputecreator.MakeHandler(s)).Methods("POST")
	r.Handle("/disputes/{dispute_id}/resolve", disputeresolver.MakeHandler(s)).Methods("PATCH")
	r.Handle("/disputes/pending", disputesfetcher.MakePendingDisputeHandler(s)).Methods("GET")
	r.Handle("/{worker_id}/disputes", disputesfetcher.MakeWorkerDisputesHandler(s)).Methods("GET")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
