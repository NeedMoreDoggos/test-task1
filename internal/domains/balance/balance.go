package balance

import "github.com/shopspring/decimal"

type Response struct {
	WalletID string
	Amount   decimal.Decimal
}

func OK(walletID string, amount decimal.Decimal) Response {
	return Response{
		WalletID: walletID,
		Amount:   amount,
	}
}

func Error(msg string) Response {
	return Response{
		WalletID: "",
		Amount:   decimal.Zero,
	}
}
