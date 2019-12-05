package disputesfetcher

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gemsorg/dispute/pkg/apierror"
	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/gemsorg/dispute/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakePendingDisputeHandler(s service.DisputeService) http.Handler {
	return kithttp.NewServer(
		makePendingDisputeFetcherEndpoint(s),
		decodePendingDisputesRequest,
		encodeResponse,
	)
}

func MakeWorkerDisputesHandler(s service.DisputeService) http.Handler {
	return kithttp.NewServer(
		makeDisputeFetcherByWorkerEndpoint(s),
		decodeWorkerDisputesRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodePendingDisputesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return dispute.Disputes{}, nil
}

func decodeWorkerDisputesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	dr := WorkerDisputesRequest{}
	vars := mux.Vars(r)
	workerID, ok := vars["worker_id"]
	if !ok {
		return nil, errorResponse(&apierror.ErrBadRouting{Param: "dispute_id"})
	}
	id, err := strconv.ParseUint(workerID, 10, 64)
	if err != nil {
		return nil, errorResponse(&apierror.ErrBadRouting{Param: "dispute_id"})
	}
	dr.WorkerID = id
	return dr, nil
}
