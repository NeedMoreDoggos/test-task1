package domains

import (
	"time"

	"github.com/shopspring/decimal"
)

type Event struct {
	From   string
	To     string
	Amount decimal.Decimal
	Time   time.Time
}
