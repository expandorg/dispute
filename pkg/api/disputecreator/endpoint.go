package disputecreator

import (
	"context"

	"github.com/expandorg/dispute/pkg/apierror"
	"github.com/expandorg/dispute/pkg/authentication"
	"github.com/expandorg/dispute/pkg/dispute"
	"github.com/expandorg/dispute/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func makeDisputecreatorEndpoint(svc service.DisputeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		req := request.(dispute.Dispute)
		saved, err := svc.CreateDispute(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
