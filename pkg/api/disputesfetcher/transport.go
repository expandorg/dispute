package disputesfetcher

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/gemsorg/dispute/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.DisputeService) http.Handler {
	return kithttp.NewServer(
		makePendingDisputeFetcherEndpoint(s),
		decodePendingDisputesRequest,
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
