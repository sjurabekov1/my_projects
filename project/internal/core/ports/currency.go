package ports

import (
	"context"

	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
)

type CurrencyService interface {
	CreateCurrency(ctx context.Context, req *pb.CreateCurrencyRequest) (*pb.CreateCurrencyResponse, error)
	DeleteCurrency(ctx context.Context, req *pb.DeleteCurrencyRequest) (*pb.DeleteCurrencyResponse, error)
	GetCurrencyList(ctx context.Context, req *pb.GetCurrencyListRequest) (*pb.GetCurrencyListResponse, error)
	GetTransitCurrencies(ctx context.Context, req *pb.GetTransitCurrenciesRequest) (*pb.GetTransitCurrenciesResponse, error)
	GetCurrencies(ctx context.Context, req *pb.GetCurrenciesRequest) (*pb.GetCurrenciesResponse, error)
}
