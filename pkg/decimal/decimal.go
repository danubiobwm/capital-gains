package decimal

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func FinancialRound(value decimal.Decimal) decimal.Decimal {
	return value.Round(2)
}

func CalculateWeightedAverage(currentQuantity int, currentAverage, newPrice decimal.Decimal, newQuantity int) decimal.Decimal {
	if currentQuantity == 0 {
		return FinancialRound(newPrice)
	}

	currentTotal := decimal.NewFromInt(int64(currentQuantity)).Mul(currentAverage)
	newTotal := decimal.NewFromInt(int64(newQuantity)).Mul(newPrice)

	totalQuantity := currentQuantity + newQuantity
	totalValue := currentTotal.Add(newTotal)

	average := totalValue.Div(decimal.NewFromInt(int64(totalQuantity)))
	return FinancialRound(average)
}

func CalculateProfit(sellPrice, weightedAverage decimal.Decimal, quantity int) decimal.Decimal {
	totalCost := decimal.NewFromInt(int64(quantity)).Mul(weightedAverage)
	totalRevenue := decimal.NewFromInt(int64(quantity)).Mul(sellPrice)
	profit := totalRevenue.Sub(totalCost)
	return FinancialRound(profit)
}

func CalculateTax(profit, accumulatedLoss decimal.Decimal) decimal.Decimal {
	if profit.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero
	}

	taxableProfit := profit
	if accumulatedLoss.GreaterThan(decimal.Zero) {
		if profit.GreaterThan(accumulatedLoss) {
			taxableProfit = profit.Sub(accumulatedLoss)
		} else {
			taxableProfit = decimal.Zero
		}
	}

	if taxableProfit.GreaterThan(decimal.Zero) {
		tax := taxableProfit.Mul(decimal.NewFromFloat(0.20))
		return FinancialRound(tax)
	}

	return decimal.Zero
}

func IsOperationExempt(unitCost decimal.Decimal, quantity int) bool {
	totalValue := unitCost.Mul(decimal.NewFromInt(int64(quantity)))
	return totalValue.LessThanOrEqual(decimal.NewFromFloat(20000.00))
}

func ParseDecimal(value string) (decimal.Decimal, error) {
	cleaned := strings.TrimSpace(value)
	cleaned = strings.ReplaceAll(cleaned, ",", ".")

	if f, err := strconv.ParseFloat(cleaned, 64); err == nil {
		return decimal.NewFromFloat(f), nil
	}

	return decimal.NewFromString(cleaned)
}

func FormatDecimal(value decimal.Decimal) string {
	return value.StringFixed(2)
}

func SumDecimals(values ...decimal.Decimal) decimal.Decimal {
	total := decimal.Zero
	for _, value := range values {
		total = total.Add(value)
	}
	return FinancialRound(total)
}

func MinDecimal(a, b decimal.Decimal) decimal.Decimal {
	if a.LessThan(b) {
		return a
	}
	return b
}

func MaxDecimal(a, b decimal.Decimal) decimal.Decimal {
	if a.GreaterThan(b) {
		return a
	}
	return b
}

func IsZero(value decimal.Decimal) bool {
	return value.Abs().LessThanOrEqual(decimal.NewFromFloat(1e-9))
}

func Percentage(value decimal.Decimal, percent float64) decimal.Decimal {
	return FinancialRound(value.Mul(decimal.NewFromFloat(percent / 100.0)))
}

func NewFromFloat(value float64) decimal.Decimal {
	// Handle special cases
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return decimal.Zero
	}
	return decimal.NewFromFloat(value)
}

func NewFromInt(value int) decimal.Decimal {
	return decimal.NewFromInt(int64(value))
}

func DivideSafe(numerator, denominator decimal.Decimal) decimal.Decimal {
	if denominator.IsZero() {
		return decimal.Zero
	}
	return FinancialRound(numerator.Div(denominator))
}

func String(value decimal.Decimal) string {
	return FormatDecimal(value)
}

func MarshalJSONForTax(value decimal.Decimal) ([]byte, error) {
	str := FormatDecimal(value)
	// Ensure it's a proper float format for JSON
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return []byte(fmt.Sprintf("%.2f", f)), nil
	}
	return []byte(str), nil
}

func CalculateWeightedAverageWithPrecision(currentQuantity int, currentAverage, newPrice decimal.Decimal, newQuantity int, precision int32) decimal.Decimal {
	if currentQuantity == 0 {
		return newPrice.Round(precision)
	}

	currentTotal := decimal.NewFromInt(int64(currentQuantity)).Mul(currentAverage)
	newTotal := decimal.NewFromInt(int64(newQuantity)).Mul(newPrice)

	totalQuantity := currentQuantity + newQuantity
	totalValue := currentTotal.Add(newTotal)

	average := totalValue.Div(decimal.NewFromInt(int64(totalQuantity)))
	return average.Round(precision)
}
