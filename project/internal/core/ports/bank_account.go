package ports

import (
	"context"

	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
)

type BankAccountService interface {
	CreateBankAccount(ctx context.Context, req *pb.CreateBankAccountRequest) (*pb.CreateBankAccountResponse, error)
	GetBankAccount(ctx context.Context, req *pb.GetBankAccountRequest) (*pb.GetBankAccountResponse, error)
	SetSalesforceId(ctx context.Context, req *pb.SetSalesforceIdRequest) (*pb.SetSalesforceIdResponse, error)
	UpdateBankAccount(ctx context.Context, req *pb.UpdateBankAccountRequest) (*pb.UpdateBankAccountResponse, error)
	UpsertBankAccountCurrency(ctx context.Context, req *pb.UpsertBankAccountCurrencyRequest) (*pb.UpsertBankAccountCurrencyResponse, error)
}
