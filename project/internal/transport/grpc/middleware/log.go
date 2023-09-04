package middleware

import (
	"context"

	"google.golang.org/grpc"

	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
)

func GrpcLoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.Log.Info("-----> Service: ", req, info, handler)
	resp, err = handler(ctx, req)

	if err != nil {
		logger.Log.Error("failed to make gRPC call: ", err)
		return resp, err
	}

	logger.Log.Info("<----- Service: ", resp)
	return resp, nil
}
