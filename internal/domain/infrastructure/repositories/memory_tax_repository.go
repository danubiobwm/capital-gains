package repositories

import (
	"github.com/danubiobwm/capital-gains/internal/domain/entities"
	"github.com/danubiobwm/capital-gains/internal/domain/repositories"
	"github.com/shopspring/decimal"
)

type MemoryTaxRepository struct {
	weightedAveragePrice decimal.Decimal
	totalStocks          int
	accumulatedLoss      decimal.Decimal
}

func NewMemoryTaxRepository() repositories.TaxRepository {
	return &MemoryTaxRepository{
		weightedAveragePrice: decimal.Zero,
		totalStocks:          0,
		accumulatedLoss:      decimal.Zero,
	}
}

func (r *MemoryTaxRepository) CalculateOperationTax(op entities.Operation) decimal.Decimal {
	switch op.Operation {
	case entities.Buy:
		return r.handleBuy(op)
	case entities.Sell:
		return r.handleSell(op)
	default:
		return decimal.Zero
	}
}

func (r *MemoryTaxRepository) handleBuy(op entities.Operation) decimal.Decimal {
	if r.totalStocks == 0 {
		r.weightedAveragePrice = op.UnitCost
		r.totalStocks = op.Quantity
	} else {
		totalValue := decimal.NewFromInt(int64(r.totalStocks)).Mul(r.weightedAveragePrice)
		newStocksValue := decimal.NewFromInt(int64(op.Quantity)).Mul(op.UnitCost)

		newTotalStocks := r.totalStocks + op.Quantity
		newTotalValue := totalValue.Add(newStocksValue)

		r.weightedAveragePrice = newTotalValue.Div(decimal.NewFromInt(int64(newTotalStocks)))
		r.totalStocks = newTotalStocks
	}

	return decimal.Zero
}

func (r *MemoryTaxRepository) handleSell(op entities.Operation) decimal.Decimal {
	totalOperationValue := op.UnitCost.Mul(decimal.NewFromInt(int64(op.Quantity)))

	// Check if total operation value is <= 20000
	if totalOperationValue.LessThanOrEqual(decimal.NewFromFloat(20000)) {
		profit := r.calculateProfit(op)
		if profit.LessThan(decimal.Zero) {
			r.accumulatedLoss = r.accumulatedLoss.Add(profit.Abs())
		}
		return decimal.Zero
	}

	profit := r.calculateProfit(op)

	// Deduct accumulated losses
	if r.accumulatedLoss.GreaterThan(decimal.Zero) {
		if profit.GreaterThan(decimal.Zero) {
			if profit.GreaterThan(r.accumulatedLoss) {
				profit = profit.Sub(r.accumulatedLoss)
				r.accumulatedLoss = decimal.Zero
			} else {
				r.accumulatedLoss = r.accumulatedLoss.Sub(profit)
				profit = decimal.Zero
			}
		} else {
			r.accumulatedLoss = r.accumulatedLoss.Add(profit.Abs())
			profit = decimal.Zero
		}
	}

	r.totalStocks -= op.Quantity

	if profit.GreaterThan(decimal.Zero) {
		tax := profit.Mul(decimal.NewFromFloat(0.20))
		return tax.RoundBank(2)
	}

	return decimal.Zero
}

func (r *MemoryTaxRepository) calculateProfit(op entities.Operation) decimal.Decimal {
	totalCost := decimal.NewFromInt(int64(op.Quantity)).Mul(r.weightedAveragePrice)
	totalRevenue := decimal.NewFromInt(int64(op.Quantity)).Mul(op.UnitCost)
	return totalRevenue.Sub(totalCost)
}

func (r *MemoryTaxRepository) ResetState() {
	r.weightedAveragePrice = decimal.Zero
	r.totalStocks = 0
	r.accumulatedLoss = decimal.Zero
}
