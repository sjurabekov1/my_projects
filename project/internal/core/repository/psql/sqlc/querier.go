// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc

import (
	"context"

	null "gopkg.in/guregu/null.v4"
)

type Querier interface {
	CreateBankAccount(ctx context.Context, arg CreateBankAccountParams) (BankAccount, error)
	CreateCompany(ctx context.Context, companyName null.String) (Company, error)
	CreateCurrency(ctx context.Context, arg CreateCurrencyParams) (Currency, error)
	CreatePocket(ctx context.Context, arg []CreatePocketParams) (int64, error)
	DeleteBankAccount(ctx context.Context, id string) error
	DeleteCompany(ctx context.Context, id string) error
	DeleteCurrency(ctx context.Context, id int32) error
	GetAllBankAccounts(ctx context.Context, arg GetAllBankAccountsParams) ([]GetAllBankAccountsRow, error)
	GetAllCompanies(ctx context.Context, arg GetAllCompaniesParams) ([]Company, error)
	GetAllCompaniesCount(ctx context.Context, companyName string) (int64, error)
	GetAllPocketTypes(ctx context.Context, arg GetAllPocketTypesParams) ([]PocketType, error)
	GetBankAccount(ctx context.Context, id string) (BankAccount, error)
	GetBankAccountById(ctx context.Context, id string) (GetBankAccountByIdRow, error)
	GetBankAccountsCount(ctx context.Context, arg GetBankAccountsCountParams) (int64, error)
	GetCompany(ctx context.Context, id string) (Company, error)
	GetCurrencies(ctx context.Context, arg GetCurrenciesParams) ([]Currency, error)
	GetCurrencyByName(ctx context.Context, name string) (Currency, error)
	GetCurrencyConvertor(ctx context.Context, arg GetCurrencyConvertorParams) ([]CurrencyConvertor, error)
	GetCurrencyConvertorByCurrency(ctx context.Context, arg GetCurrencyConvertorByCurrencyParams) (CurrencyConvertor, error)
	GetPocketsByUserID(ctx context.Context, userID null.String) ([]GetPocketsByUserIDRow, error)
	SetSalesforceId(ctx context.Context, arg SetSalesforceIdParams) error
	SoftDeleteCompany(ctx context.Context, id string) error
	UpdateBankAccount(ctx context.Context, arg UpdateBankAccountParams) error
	UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error
	UpdateCurrency(ctx context.Context, arg UpdateCurrencyParams) (Currency, error)
	UpdateCurrencyConvertor(ctx context.Context, arg UpdateCurrencyConvertorParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
