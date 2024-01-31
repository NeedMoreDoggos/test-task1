package transfermoney

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/NeedMoreDoggos/test-task1/internal/domains/transfer"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/shopspring/decimal"
)

var (
	ErrEmptyTo         = errors.New("empty to")
	ErrEmptyAmount     = errors.New("empty amount")
	ErrUncorrectAmount = errors.New("incorrect amount")
)

type accToAccTransferer interface {
	Transfer(from, to string, amount decimal.Decimal) error
}

func New(log *slog.Logger, accToAccTransferer accToAccTransferer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.TransferMoney"
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		from := chi.URLParam(r, "walletid")

		var req transfer.Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error(err.Error())
			render.JSON(w, r, transfer.Error("failed to parse request"))
			return
		}

		if err := validateReq(req); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, transfer.Error("failed to validate request"))
			return
		}

		amount, err := decimal.NewFromString(req.Amount)
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, transfer.Error("failed to parse amount"))
			return
		}

		err = accToAccTransferer.Transfer(from, req.To, amount)
		if err != nil {
			if errors.Is(err, errors.New("outgoing wallet not found")) {
				log.Error(err.Error())
				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, transfer.Error("outgoing wallet not found"))
				return
			}

			if errors.Is(err, errors.New("insufficient funds")) || errors.Is(err, errors.New("negative balance")) {
				log.Error(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, transfer.Error("insufficient funds"))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, transfer.OK())

		log.Info("wallet transfered", slog.String("wallet_id", from),
			slog.String("to_wallet_id", req.To),
			slog.String("amount", amount.String()))
	}
}

func validateReq(r transfer.Request) error {
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
