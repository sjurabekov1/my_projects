package service

import (
	"context"

	"hyssa/hyssa_go_billing_service/internal/core/ports"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/pkg/processor"
	"hyssa/hyssa_go_billing_service/internal/pkg/serialize"

	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
)

var (
	_ ports.BankAccountService = (*BankAccount)(nil)
)

type BankAccount struct {
	pb.UnimplementedBankServiceServer
	repo repository.Store
}

func NewBankAccountService(repo repository.Store) *BankAccount {
	return &BankAccount{
		repo: repo,
	}
}

func (b *BankAccount) CreateBankAccount(ctx context.Context, req *pb.CreateBankAccountRequest) (resp *pb.CreateBankAccountResponse, err error) {
	bankAccountDB, err := processor.ExecuteWithResp(ctx, req, b.repo.CreateBankAccount)
	if err != nil {
		return nil, err
	}

	err = serialize.MarshalUnMarshal(ctx, bankAccountDB, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BankAccount) GetBankAccount(ctx context.Context, req *pb.GetBankAccountRequest) (resp *pb.GetBankAccountResponse, err error) {
	bankAccountDB, err := processor.ExecuteWithResp(ctx, req.GetId(), b.repo.GetBankAccount)
	if err != nil {
		return nil, err
	}

	err = serialize.MarshalUnMarshal(ctx, bankAccountDB, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BankAccount) SetSalesforceId(ctx context.Context, req *pb.SetSalesforceIdRequest) (*pb.SetSalesforceIdResponse, error) {
	// Ask what is the purpose of this function
	return &pb.SetSalesforceIdResponse{}, nil
}

func (b *BankAccount) UpdateBankAccount(ctx context.Context, req *pb.UpdateBankAccountRequest) (resp *pb.UpdateBankAccountResponse, err error) {
	err = processor.Execute(ctx, req, b.repo.UpdateBankAccount)
	return resp, err
}

func (b *BankAccount) UpsertBankAccountCurrency(ctx context.Context, req *pb.UpsertBankAccountCurrencyRequest) (
	*pb.UpsertBankAccountCurrencyResponse, error) {
	// Ask what is the purpose of this function
	return &pb.UpsertBankAccountCurrencyResponse{}, nil
}
