package transfer

import (
	"errors"

	"github.com/shopspring/decimal"
)

var (
	ErrEmptyTo         = errors.New("empty to")
	ErrEmptyAmount     = errors.New("empty amount")
	ErrUncorrectAmount = errors.New("incorrect amount")
)

type Request struct {
	To            string `json:"to"`
	Amount        string `json:"amount"`
	amountDecimal *decimal.Decimal
}

func ValidateReq(r Request) error {
	if r.To == "" {
		return ErrEmptyTo
	}

	if r.Amount == "" {
		return ErrEmptyAmount
	}

	_, err := decimal.NewFromString(r.Amount)
	if err != nil {
		return ErrUncorrectAmount
	}

	return nil
}
