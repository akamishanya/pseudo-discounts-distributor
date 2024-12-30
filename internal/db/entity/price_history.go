package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type PriceHistory struct {
	Id        int64
	ProductId int64
	Price     decimal.Decimal
	date      time.Time
}
