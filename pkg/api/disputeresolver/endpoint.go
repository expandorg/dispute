package disputeresolver

import (
	"context"

	"github.com/expandorg/dispute/pkg/apierror"
	"github.com/expandorg/dispute/pkg/authentication"
	"github.com/expandorg/dispute/pkg/dispute"
	"github.com/expandorg/dispute/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func makeDisputeResolverEndpoint(svc service.DisputeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		req := request.(dispute.Resolution)
		resolved, err := svc.ResolveDispute(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return resolved, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(401, err.Error(), err)
}
