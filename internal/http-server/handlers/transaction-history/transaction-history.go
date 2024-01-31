package transactionhistory

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NeedMoreDoggos/test-task1/internal/domains"
	"github.com/NeedMoreDoggos/test-task1/internal/domains/historyan"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type history interface {
	//Тут есть возможная проблема, что если транзакций будет очень много
	//То будет не нужная работа + на фрон придет много лишнего
	//Так что есть куда улучшать.
	//Буду считать, что достаточно будет последние 20 транзакции
	//Это можно поправить все с помощью стриминга
	Transactions(walletID string) ([]domains.Event, error)
}

func New(log *slog.Logger, h history) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.TransactionHistory"
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		walletID := chi.URLParam(r, "walletid")

		events, err := h.Transactions(walletID)
		if err != nil {
			if errors.Is(err, fmt.Errorf("wallet %s not found", walletID)) {
				log.Error(err.Error())
				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, historyan.Error("failed to get transactions"))
				return
			}

			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, historyan.Error("internal server error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, historyan.OK(events))

		log.Info("transactions fetched", slog.String("wallet_id", walletID))
	}
}
