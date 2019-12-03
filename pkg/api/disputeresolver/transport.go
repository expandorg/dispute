package disputeresolver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/dispute/pkg/apierror"
	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/gemsorg/dispute/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.DisputeService) http.Handler {
	return kithttp.NewServer(
		makeDisputeResolverEndpoint(s),
		decodeDisputeResolverRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeDisputeResolverRequest(_ context.Context, r *http.Request) (interface{}, error) {
	dr := dispute.Resolution{}
	vars := mux.Vars(r)
	disputeID, ok := vars["dispute_id"]
	dr.DisputeID = disputeID
	if !ok {
		return nil, errorResponse(&apierror.ErrBadRouting{Param: "dispute_id"})
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dr)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	return dr, nil
}
