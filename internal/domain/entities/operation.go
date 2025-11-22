package entities

import "github.com/shopspring/decimal"

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Operation struct {
	Operation OperationType   `json:"operation"`
	UnitCost  decimal.Decimal `json:"unit-cost"`
	Quantity  int             `json:"quantity"`
}

type TaxResult struct {
	Tax decimal.Decimal `json:"tax"`
}
