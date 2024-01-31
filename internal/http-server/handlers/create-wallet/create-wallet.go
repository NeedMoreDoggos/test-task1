package createwallet

import (
	"log/slog"
	"net/http"

	creatres "github.com/NeedMoreDoggos/test-task1/internal/domains/create"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/shopspring/decimal"
)

type walletCreater interface {
	CreateWallet() (string, decimal.Decimal, error)
}

func New(log *slog.Logger, wawalletCreater walletCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.CreateWallet"
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id, balance, err := wawalletCreater.CreateWallet()
		if err != nil {
			log.Error(err.Error())
			render.JSON(w, r, creatres.Error("failed to create wallet"))
			return
		}

		log.Info("wallet created", slog.String("wallet_id", id))

		render.JSON(w, r, creatres.OK(id, balance.String()))
	}
}
