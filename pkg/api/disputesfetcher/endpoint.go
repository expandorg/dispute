package disputesfetcher

import (
	"context"

	"github.com/expandorg/dispute/pkg/apierror"
	"github.com/expandorg/dispute/pkg/authentication"
	"github.com/expandorg/dispute/pkg/dispute"
	"github.com/expandorg/dispute/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func makePendingDisputeFetcherEndpoint(svc service.DisputeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		disp, err := svc.GetDisputesByStatus(dispute.Pending)
		if err != nil {
			return nil, errorResponse(err)
		}
		return DisputesResponse{disp}, nil
	}
}

func makeDisputeFetcherByWorkerEndpoint(svc service.DisputeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		req := request.(WorkerDisputesRequest)

		disp, err := svc.GetDisputesByWorkerID(req.WorkerID)
		if err != nil {
			return nil, errorResponse(err)
		}
		return DisputesResponse{disp}, nil
	}
}

type WorkerDisputesRequest struct {
	WorkerID uint64
}

type DisputesResponse struct {
	Disputes dispute.Disputes `json:"disputes"`
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
