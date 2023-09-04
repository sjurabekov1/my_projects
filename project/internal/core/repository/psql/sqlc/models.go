// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	null "gopkg.in/guregu/null.v4"
)

type TaxType string

const (
	TaxTypeFixed      TaxType = "Fixed"
	TaxTypePercentage TaxType = "Percentage"
)

func (e *TaxType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TaxType(s)
	case string:
		*e = TaxType(s)
	default:
		return fmt.Errorf("unsupported scan type for TaxType: %T", src)
	}
	return nil
}

type NullTaxType struct {
	TaxType TaxType
	Valid   bool // Valid is true if TaxType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaxType) Scan(value interface{}) error {
	if value == nil {
		ns.TaxType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TaxType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaxType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TaxType), nil
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusCanceled  TransactionStatus = "canceled"
)

func (e *TransactionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionStatus(s)
	case string:
		*e = TransactionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionStatus: %T", src)
	}
	return nil
}

type NullTransactionStatus struct {
	TransactionStatus TransactionStatus
	Valid             bool // Valid is true if TransactionStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionStatus), nil
}

type BankAccount struct {
	ID                        string      `json:"id"`
	OID                       int32       `json:"o_id"`
	BankAccountName           null.String `json:"bank_account_name"`
	BicCode                   null.String `json:"bic_code"`
	IbanCode                  null.String `json:"iban_code"`
	SwiftCode                 null.String `json:"swift_code"`
	Region                    null.String `json:"region"`
	RegistrationNumber        null.Int    `json:"registration_number"`
	Address                   null.String `json:"address"`
	BankAddress               null.String `json:"bank_address"`
	PhoneNumber               null.String `json:"phone_number"`
	PaymentReference          null.String `json:"payment_reference"`
	Description               null.String `json:"description"`
	SalesforceID              null.String `json:"salesforce_id"`
	OwnerName                 null.String `json:"owner_name"`
	TransactionsReference     null.String `json:"transactions_reference"`
	ReceivingPartyBankDetails null.String `json:"receiving_party_bank_details"`
	Currency                  int32       `json:"currency"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
	DeletedAt                 null.Time   `json:"deleted_at"`
}

type BankAccountCurrency struct {
	ID            int32        `json:"id"`
	BankAccountID null.String  `json:"bank_account_id"`
	CurrencyID    null.Int     `json:"currency_id"`
	CreatedAt     sql.NullTime `json:"created_at"`
}

type Broker struct {
	ID            string      `json:"id"`
	OID           int32       `json:"o_id"`
	Name          null.String `json:"name"`
	LegalAddrress null.String `json:"legal_addrress"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     null.Time   `json:"deleted_at"`
}

type ComissionGroup struct {
	ID   string      `json:"id"`
	OID  int32       `json:"o_id"`
	Name null.String `json:"name"`
}

type ComissionGroupTaxRule struct {
	ID               string      `json:"id"`
	OID              int32       `json:"o_id"`
	ComissionGroupID null.String `json:"comission_group_id"`
	TaxID            null.String `json:"tax_id"`
}

type Company struct {
	ID          string      `json:"id"`
	OID         int32       `json:"o_id"`
	CompanyName null.String `json:"company_name"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   null.Time   `json:"deleted_at"`
}

type Currency struct {
	ID     int32       `json:"id"`
	Name   string      `json:"name"`
	Icon   null.String `json:"icon"`
	Active null.Bool   `json:"active"`
}

type CurrencyConvertor struct {
	ID                int32      `json:"id"`
	PrimaryCurrency   int32      `json:"primary_currency"`
	SecondaryCurrency int32      `json:"secondary_currency"`
	Amount            null.Float `json:"amount"`
}

type MainGraph struct {
	InvestorID null.String `json:"investor_id"`
	Taum       null.String `json:"taum"`
	PocketID   null.String `json:"pocket_id"`
	SymbolID   null.String `json:"symbol_id"`
}

type Operation struct {
	ID            string      `json:"id"`
	OperationName null.String `json:"operation_name"`
}

type Order struct {
	ID                    string      `json:"id"`
	OID                   int64       `json:"o_id"`
	InvestorID            null.String `json:"investor_id"`
	SymbolID              null.String `json:"symbol_id"`
	PlaceTime             null.Time   `json:"place_time"`
	CurrentModificationID null.String `json:"current_modification_id"`
	ExanteAccountID       null.String `json:"exante_account_id"`
	Username              null.String `json:"username"`
	Status                null.String `json:"status"`
	LastUpdate            null.String `json:"last_update"`
	Reason                null.String `json:"reason"`
	LimitPrice            null.String `json:"limit_price"`
	Quantity              null.String `json:"quantity"`
	Side                  null.String `json:"side"`
	SalesforceOrderID     null.String `json:"salesforce_order_id"`
	PocketID              null.String `json:"pocket_id"`
	Currency              null.String `json:"currency"`
	FuturePocketID        null.String `json:"future_pocket_id"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
	DeletedAt             null.Time   `json:"deleted_at"`
}

type Partner struct {
	ID        string      `json:"id"`
	OID       int32       `json:"o_id"`
	Name      null.String `json:"name"`
	UserType  null.String `json:"user_type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt null.Time   `json:"deleted_at"`
}

type Pocket struct {
	ID           string       `json:"id"`
	OID          int64        `json:"o_id"`
	Name         null.String  `json:"name"`
	UserID       null.String  `json:"user_id"`
	CompanyID    null.String  `json:"company_id"`
	Type         null.Int     `json:"type"`
	CreatedAt    sql.NullTime `json:"created_at"`
	Title        null.String  `json:"title"`
	Icon         null.String  `json:"icon"`
	SalesforceID null.String  `json:"salesforce_id"`
	Color        null.String  `json:"color"`
}

type PocketType struct {
	ID       int32       `json:"id"`
	Name     null.String `json:"name"`
	IconID   null.String `json:"icon_id"`
	Required null.Bool   `json:"required"`
}

type Tax struct {
	ID                string         `json:"id"`
	OID               int32          `json:"o_id"`
	Type              NullTaxType    `json:"type"`
	Value             pgtype.Numeric `json:"value"`
	OperationID       null.String    `json:"operation_id"`
	CompanyID         null.String    `json:"company_id"`
	Currency          null.String    `json:"currency"`
	MinFee            pgtype.Numeric `json:"min_fee"`
	BrokerID          null.String    `json:"broker_id"`
	PartnerID         null.String    `json:"partner_id"`
	Discount          pgtype.Numeric `json:"discount"`
	BrokerType        NullTaxType    `json:"broker_type"`
	BrokerCommission  pgtype.Numeric `json:"broker_commission"`
	PartnerCommission pgtype.Numeric `json:"partner_commission"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         null.Time      `json:"deleted_at"`
}

type Transaction struct {
	ID             string                `json:"id"`
	ParentID       null.String           `json:"parent_id"`
	OID            int64                 `json:"o_id"`
	Comment        null.String           `json:"comment"`
	UserID         null.String           `json:"user_id"`
	PocketID       null.String           `json:"pocket_id"`
	OperationID    null.String           `json:"operation_id"`
	Amount         pgtype.Numeric        `json:"amount"`
	Currency       null.String           `json:"currency"`
	Status         NullTransactionStatus `json:"status"`
	OrderID        null.String           `json:"order_id"`
	BankAccount    null.String           `json:"bank_account"`
	SalesforceID   null.String           `json:"salesforce_id"`
	FuturePocketID null.String           `json:"future_pocket_id"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	DeletedAt      null.Time             `json:"deleted_at"`
}