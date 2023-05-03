package server_lib

import (
	"context"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandlePanic() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandlerContext(
		func(ctx context.Context, p interface{}) (err error) {
			return status.Errorf(codes.Internal, "%s", p)
		})
}
