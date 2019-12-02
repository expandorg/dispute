package healthchecker

import (
	"context"

	service "github.com/gemsorg/dispute/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeHealthyEndpoint(svc service.DisputeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		healthy := svc.Healthy()
		return HealthyResponse{healthy}, nil
	}
}

type HealthyResponse struct {
	Healthy bool `json:"healthy"`
}
