package balance

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NeedMoreDoggos/test-task1/internal/domains/balance"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/shopspring/decimal"
)

type balanceRetriever interface {
	Balance(walletID string) (decimal.Decimal, error)
}

func New(log *slog.Logger, b balanceRetriever) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.Balance"
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		walletID := chi.URLParam(r, "walletid")

		bal, err := b.Balance(walletID)
		if err != nil {
			if errors.Is(err, fmt.Errorf("wallet %s not found", walletID)) {
				log.Error(err.Error())
				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, balance.Error("wallet not found"))
				return
			}

			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, balance.Error("internal server error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, balance.OK(walletID, bal))
	}
}
