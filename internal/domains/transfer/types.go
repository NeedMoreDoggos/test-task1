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
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	OK    bool   `json:"ok"`
}

func Error(msg string) Response {
	return Response{
		OK:    false,
		Error: msg,
	}
}

func OK() Response {
	return Response{
		OK: true,
	}
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
