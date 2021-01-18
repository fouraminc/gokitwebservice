package gokitwebservice

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)
type Endpoints struct {
	GetEndPoint endpoint.Endpoint
	StatusEndPoint endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getRequest)
		d, err := srv.Get(ctx)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return	getResponse{d,""}, nil
	}
}
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest) // we really just need the request, we don't use any value from it
		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		b, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return validateResponse{b, err.Error()}, nil
		}
		return validateResponse{b, ""}, nil
	}
}

func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := getRequest{}
	resp, err := e.GetEndPoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}


func (e Endpoints) Status (ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndPoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

func (e Endpoints) Validate(ctx context.Context, date string) (bool, error){
	req := validateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(validateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}