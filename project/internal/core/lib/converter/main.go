package converter

import (
	"context"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"

	"github.com/go-playground/validator/v10"
)

type ConvertCurrencyRequest struct {
	FromCurrency string  `validate:"required,max=3"`
	ToCurrency   string  `validate:"required,max=3"`
	Amount       float64 `validate:"required"`
	Ratio        float64 `validate:"required"`
}

func (r *ConvertCurrencyRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func ConvertCurrency(ctx context.Context, req *ConvertCurrencyRequest) (resp float64, err error) {
	err = req.Validate()
	if err != nil {
		return 0, err
	}

	logger.Log.Debug("ConvertCurrency request from to amount ratio: ", req.FromCurrency, req.ToCurrency, req.Amount, req.Ratio)

	resp = req.Amount * req.Ratio
	return
}
