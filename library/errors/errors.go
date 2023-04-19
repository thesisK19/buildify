package errors_lib

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToDefaultError(err error) error {
	if err == nil {
		return status.Error(codes.Internal, codes.Internal.String())
	}
	return status.Error(codes.Internal, err.Error())
}

func ToNotFoundError(err error) error {
	if err == nil {
		return status.Error(codes.NotFound, codes.NotFound.String())
	}
	return status.Error(codes.NotFound, err.Error())
}

func ToUnauthenticatedError(err error) error {
	if err == nil {
		return status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	return status.Error(codes.Unauthenticated, err.Error())
}

func ToInvalidArgumentError(err error) error {
	if err == nil {
		return status.Error(codes.InvalidArgument, codes.Internal.String())
	}
	return status.Error(codes.InvalidArgument, err.Error())
}

func GetStatusCodeFromError(err error) codes.Code {
	status, _ := status.FromError(err)
	return status.Code()
}
