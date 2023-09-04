package service

import (
	"context"
	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
	"hyssa/hyssa_go_billing_service/internal/core/ports"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/core/repository/psql/sqlc"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
	"hyssa/hyssa_go_billing_service/internal/pkg/processor"
	"hyssa/hyssa_go_billing_service/internal/pkg/serialize"
)

var (
	_ ports.CurrencyService = (*Currency)(nil)
)

type Currency struct {
	pb.UnimplementedCurrencyServiceServer
	repo repository.Store
}

func NewCurrencyService(repo repository.Store) *Currency {
	return &Currency{
		repo: repo,
	}
}

func (c *Currency) CreateCurrency(ctx context.Context, req *pb.CreateCurrencyRequest) (resp *pb.CreateCurrencyResponse, err error) {
	var (
		dbResp sqlc.Currency
	)

	if req.Id == 0 {
		dbResp, err = processor.ExecuteWithResp(ctx, req, c.repo.CreateCurrency)
	} else {
		dbResp, err = processor.ExecuteWithResp(ctx, req, c.repo.UpdateCurrency)
	}

	if err != nil {
		return nil, err
	}

	err = serialize.MarshalUnMarshal(ctx, dbResp, &resp)
	if err != nil {
		return nil, err
	}

	return
}

func (c *Currency) DeleteCurrency(ctx context.Context, req *pb.DeleteCurrencyRequest) (resp *pb.DeleteCurrencyResponse, err error) {
	err = processor.Execute(ctx, req.Id, c.repo.DeleteCurrency)
	return
}

func (c *Currency) GetCurrencyList(ctx context.Context, req *pb.GetCurrencyListRequest) (resp *pb.GetCurrencyListResponse, err error) {
	logger.Log.Error("not implemented yet")
	// depends on pocket
	return
}

func (c *Currency) GetTransitCurrencies(ctx context.Context, req *pb.GetTransitCurrenciesRequest) (resp *pb.GetTransitCurrenciesResponse, err error) {
	// depends on pocket
	logger.Log.Error("not implemented yet")
	return
}

func (c *Currency) GetCurrencies(ctx context.Context, req *pb.GetCurrenciesRequest) (resp *pb.GetCurrenciesResponse, err error) {
	dbResp, err := processor.ExecuteWithResp(ctx, req, c.repo.GetCurrencies)
	if err != nil {
		return
	}

	resp = &pb.GetCurrenciesResponse{}
	err = serialize.MarshalUnMarshal(ctx, dbResp, &resp.Currencies)
	if err != nil {
		return
	}

	return
}
