package disputeresolver

import (
	"context"

	"github.com/gemsorg/dispute/pkg/apierror"
	"github.com/gemsorg/dispute/pkg/authentication"
	"github.com/gemsorg/dispute/pkg/dispute"
	"github.com/gemsorg/dispute/pkg/service"
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
