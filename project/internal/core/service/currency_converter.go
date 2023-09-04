package service

import (
	"context"
	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
	"hyssa/hyssa_go_billing_service/internal/core/lib/converter"
	"hyssa/hyssa_go_billing_service/internal/core/ports"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/core/repository/psql/sqlc"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
	"hyssa/hyssa_go_billing_service/internal/pkg/processor"
	"hyssa/hyssa_go_billing_service/internal/pkg/serialize"
)

var (
	_ ports.CurrencyConverterService = (*CurrencyConverter)(nil)
)

type CurrencyConverter struct {
	pb.UnimplementedCurrencyConverterServiceServer
	repo repository.Store
}

func NewCurrencyConverterService(repo repository.Store) *CurrencyConverter {
	return &CurrencyConverter{
		repo: repo,
	}
}

func (c *CurrencyConverter) CreateCurrencyConvertor(ctx context.Context, req *pb.CreateCurrencyConvertorRequest) (
	resp *pb.CreateCurrencyConvertorResponse, err error) {
	dbResp, err := processor.ExecuteWithResp(ctx, req, c.repo.UpdateCurrencyConvertor)
	if err != nil {
		return nil, err
	}

	resp = &pb.CreateCurrencyConvertorResponse{}
	err = serialize.MarshalUnMarshal(ctx, dbResp, &resp.Id)
	if err != nil {
		return nil, err
	}

	return
}

func (c *CurrencyConverter) GetCurrencyConvertor(ctx context.Context, req *pb.GetCurrencyConvertorRequest) (
	resp *pb.GetCurrencyConvertorResponse, err error) {
	logger.Log.Error("not implemented yet")
	return
}

func (c *CurrencyConverter) UpdateCurrencyConvertor(ctx context.Context, req *pb.UpdateCurrencyConvertorRequest) (
	resp *pb.UpdateCurrencyConvertorResponse, err error) {
	logger.Log.Error("not implemented yet")
	return
}

func (c *CurrencyConverter) DeleteCurrencyConvertor(ctx context.Context, req *pb.DeleteCurrencyConvertorRequest) (
	resp *pb.DeleteCurrencyConvertorResponse, er error) {
	logger.Log.Error("not implemented yet")
	return
}

func (c *CurrencyConverter) GetNameBySalesforceID(ctx context.Context, req *pb.GetNameBySalesforceIDRequest) (
	resp *pb.GetNameBySalesforceIDResponse, err error) {
	// TODO: Need to ask why do we need this function and where it is used
	logger.Log.Error("not implemented yet")
	resp = &pb.GetNameBySalesforceIDResponse{}
	return
}

func (c *CurrencyConverter) ConvertCurrency(ctx context.Context, req *pb.ConvertCurrencyRequest) (
	resp *pb.ConvertCurrencyResponse, err error) {
	fromCurrency, err := c.repo.GetCurrencyByName(ctx, req.FromCurrency)
	if err != nil {
		return nil, err
	}

	toCurrency, err := c.repo.GetCurrencyByName(ctx, req.ToCurrency)
	if err != nil {
		return nil, err
	}

	convertorItem, err := c.repo.GetCurrencyConvertorByCurrency(ctx, sqlc.GetCurrencyConvertorByCurrencyParams{
		PrimaryCurrency:   fromCurrency.ID,
		SecondaryCurrency: toCurrency.ID,
	})
	if err != nil {
		return nil, err
	}

	convertedAmount, err := converter.ConvertCurrency(ctx, &converter.ConvertCurrencyRequest{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Amount:       req.Amount,
		Ratio:        convertorItem.Amount.Float64,
	})

	if err != nil {
		return nil, err
	}

	resp = &pb.ConvertCurrencyResponse{
		Amount:   convertedAmount,
		Currency: req.GetToCurrency(),
	}

	return
}
