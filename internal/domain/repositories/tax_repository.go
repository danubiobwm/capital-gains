package repositories

import (
	"github.com/danubiobwm/capital-gains/internal/domain/entities"
	"github.com/shopspring/decimal"
)

type TaxRepository interface {
	CalculateOperationTax(operation entities.Operation) decimal.Decimal
	ResetState()
}
