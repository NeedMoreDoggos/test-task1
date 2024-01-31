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
