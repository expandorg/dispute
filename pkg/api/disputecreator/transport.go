package disputecreator

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/dispute/pkg/apierror"
	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/gemsorg/dispute/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.DisputeService) http.Handler {
	return kithttp.NewServer(
		makeDisputecreatorEndpoint(s),
		decodeNewDisputeRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeNewDisputeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var d dispute.Dispute
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	return d, nil
}
