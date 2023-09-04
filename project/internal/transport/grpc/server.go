package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/core/service"
	"hyssa/hyssa_go_billing_service/internal/transport/grpc/middleware"
)

func New(repo repository.Store) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.GrpcLoggerMiddleware,
			),
		),
	)

	reflection.Register(grpcServer)

	billing_service_v2.RegisterBankServiceServer(grpcServer, service.NewBankAccountService(repo))
	billing_service_v2.RegisterCurrencyServiceServer(grpcServer, service.NewCurrencyService(repo))
	billing_service_v2.RegisterPocketServiceServer(grpcServer, service.NewPocketService(repo))

	return grpcServer
}
