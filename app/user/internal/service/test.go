package service

import (
	"context"
	"fmt"

	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/library/errors"
)

func (s *Service) Test(ctx context.Context, in *api.TestRequest) (*api.TestResponse, error) {
	// logger := ctxlogrus.Extract(ctx).WithField("func", "Test")

	err := fmt.Errorf("alo alo errr")
	if in.Id == 1 {
		return nil, errors.ToDefaultError(err)
	}
	if in.Id == 2 {
		return nil, errors.ToInvalidArgumentError(err)
	}
	if in.Id == 3 {
		return nil, errors.ToNotFoundError(err)
	}

	return nil, err
}
