package ports

import (
	"context"

	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
)

type CurrencyConverterService interface {
	CreateCurrencyConvertor(ctx context.Context, req *pb.CreateCurrencyConvertorRequest) (*pb.CreateCurrencyConvertorResponse, error)
	GetCurrencyConvertor(ctx context.Context, req *pb.GetCurrencyConvertorRequest) (*pb.GetCurrencyConvertorResponse, error)
	UpdateCurrencyConvertor(ctx context.Context, req *pb.UpdateCurrencyConvertorRequest) (*pb.UpdateCurrencyConvertorResponse, error)
	DeleteCurrencyConvertor(ctx context.Context, req *pb.DeleteCurrencyConvertorRequest) (*pb.DeleteCurrencyConvertorResponse, error)
	GetNameBySalesforceID(ctx context.Context, req *pb.GetNameBySalesforceIDRequest) (*pb.GetNameBySalesforceIDResponse, error)
	ConvertCurrency(ctx context.Context, req *pb.ConvertCurrencyRequest) (*pb.ConvertCurrencyResponse, error)
}
