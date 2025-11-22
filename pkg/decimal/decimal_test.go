package decimal

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFinancialRound(t *testing.T) {
	tests := []struct {
		input    decimal.Decimal
		expected decimal.Decimal
	}{
		{decimal.NewFromFloat(10.123), decimal.NewFromFloat(10.12)},
		{decimal.NewFromFloat(10.125), decimal.NewFromFloat(10.13)},
		{decimal.NewFromFloat(10.124), decimal.NewFromFloat(10.12)},
		{decimal.NewFromFloat(0.001), decimal.NewFromFloat(0.00)},
		{decimal.NewFromFloat(999.999), decimal.NewFromFloat(1000.00)},
		{decimal.NewFromFloat(10.135), decimal.NewFromFloat(10.14)}, // Additional test case
	}

	for _, tt := range tests {
		t.Run(tt.input.String(), func(t *testing.T) {
			result := FinancialRound(tt.input)
			assert.True(t, result.Equal(tt.expected),
				"Input: %s, Expected %s, got %s", tt.input.String(), tt.expected.String(), result.String())
		})
	}
}

func TestCalculateWeightedAverage(t *testing.T) {
	tests := []struct {
		name            string
		currentQty      int
		currentAvg      decimal.Decimal
		newPrice        decimal.Decimal
		newQty          int
		expectedAverage decimal.Decimal
	}{
		{
			name:            "First purchase",
			currentQty:      0,
			currentAvg:      decimal.Zero,
			newPrice:        decimal.NewFromFloat(10.00),
			newQty:          100,
			expectedAverage: decimal.NewFromFloat(10.00),
		},
		{
			name:            "Additional purchase same price",
			currentQty:      100,
			currentAvg:      decimal.NewFromFloat(10.00),
			newPrice:        decimal.NewFromFloat(10.00),
			newQty:          50,
			expectedAverage: decimal.NewFromFloat(10.00),
		},
		{
			name:            "Additional purchase different price",
			currentQty:      100,
			currentAvg:      decimal.NewFromFloat(10.00),
			newPrice:        decimal.NewFromFloat(20.00),
			newQty:          50,
			expectedAverage: decimal.NewFromFloat(13.33),
		},
		{
			name:            "Complex weighted average",
			currentQty:      5,
			currentAvg:      decimal.NewFromFloat(20.00),
			newPrice:        decimal.NewFromFloat(10.00),
			newQty:          5,
			expectedAverage: decimal.NewFromFloat(15.00),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateWeightedAverage(tt.currentQty, tt.currentAvg, tt.newPrice, tt.newQty)
			assert.True(t, result.Equal(tt.expectedAverage),
				"Expected %s, got %s", tt.expectedAverage.String(), result.String())
		})
	}
}

func TestCalculateProfit(t *testing.T) {
	tests := []struct {
		name           string
		sellPrice      decimal.Decimal
		weightedAvg    decimal.Decimal
		quantity       int
		expectedProfit decimal.Decimal
	}{
		{
			name:           "Profit scenario",
			sellPrice:      decimal.NewFromFloat(20.00),
			weightedAvg:    decimal.NewFromFloat(10.00),
			quantity:       100,
			expectedProfit: decimal.NewFromFloat(1000.00),
		},
		{
			name:           "Loss scenario",
			sellPrice:      decimal.NewFromFloat(5.00),
			weightedAvg:    decimal.NewFromFloat(10.00),
			quantity:       100,
			expectedProfit: decimal.NewFromFloat(-500.00),
		},
		{
			name:           "Break even",
			sellPrice:      decimal.NewFromFloat(15.00),
			weightedAvg:    decimal.NewFromFloat(15.00),
			quantity:       100,
			expectedProfit: decimal.Zero,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateProfit(tt.sellPrice, tt.weightedAvg, tt.quantity)
			assert.True(t, result.Equal(tt.expectedProfit),
				"Expected %s, got %s", tt.expectedProfit.String(), result.String())
		})
	}
}

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name            string
		profit          decimal.Decimal
		accumulatedLoss decimal.Decimal
		expectedTax     decimal.Decimal
	}{
		{
			name:            "Profit with no losses",
			profit:          decimal.NewFromFloat(1000.00),
			accumulatedLoss: decimal.Zero,
			expectedTax:     decimal.NewFromFloat(200.00),
		},
		{
			name:            "Profit with partial loss deduction",
			profit:          decimal.NewFromFloat(1000.00),
			accumulatedLoss: decimal.NewFromFloat(500.00),
			expectedTax:     decimal.NewFromFloat(100.00),
		},
		{
			name:            "Profit completely offset by losses",
			profit:          decimal.NewFromFloat(1000.00),
			accumulatedLoss: decimal.NewFromFloat(1500.00),
			expectedTax:     decimal.Zero,
		},
		{
			name:            "No profit",
			profit:          decimal.NewFromFloat(-500.00),
			accumulatedLoss: decimal.Zero,
			expectedTax:     decimal.Zero,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateTax(tt.profit, tt.accumulatedLoss)
			assert.True(t, result.Equal(tt.expectedTax),
				"Expected %s, got %s", tt.expectedTax.String(), result.String())
		})
	}
}

func TestIsOperationExempt(t *testing.T) {
	tests := []struct {
		name     string
		unitCost decimal.Decimal
		quantity int
		expected bool
	}{
		{
			name:     "Exempt - below threshold",
			unitCost: decimal.NewFromFloat(10.00),
			quantity: 1000, // 10 * 1000 = 10000
			expected: true,
		},
		{
			name:     "Exempt - exactly at threshold",
			unitCost: decimal.NewFromFloat(20.00),
			quantity: 1000, // 20 * 1000 = 20000
			expected: true,
		},
		{
			name:     "Not exempt - above threshold",
			unitCost: decimal.NewFromFloat(20.01),
			quantity: 1000, // 20.01 * 1000 = 20010
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOperationExempt(tt.unitCost, tt.quantity)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected decimal.Decimal
		hasError bool
	}{
		{"10.50", decimal.NewFromFloat(10.50), false},
		{"10,50", decimal.NewFromFloat(10.50), false},
		{"1000", decimal.NewFromFloat(1000.00), false},
		{"abc", decimal.Zero, true},
		{"", decimal.Zero, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDecimal(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, result.Equal(tt.expected))
			}
		})
	}
}

func TestSumDecimals(t *testing.T) {
	values := []decimal.Decimal{
		decimal.NewFromFloat(10.50),
		decimal.NewFromFloat(20.25),
		decimal.NewFromFloat(5.75),
	}

	result := SumDecimals(values...)
	expected := decimal.NewFromFloat(36.50)

	assert.True(t, result.Equal(expected))
}

func TestMinMaxDecimal(t *testing.T) {
	a := decimal.NewFromFloat(10.50)
	b := decimal.NewFromFloat(20.25)

	assert.True(t, MinDecimal(a, b).Equal(a))
	assert.True(t, MaxDecimal(a, b).Equal(b))
}

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero(decimal.Zero))
	assert.True(t, IsZero(decimal.NewFromFloat(0.000000001))) // Within tolerance
	assert.False(t, IsZero(decimal.NewFromFloat(0.01)))
	assert.True(t, IsZero(decimal.NewFromFloat(0.0000000001))) // Even smaller
}

func TestPercentage(t *testing.T) {
	value := decimal.NewFromFloat(1000.00)
	result := Percentage(value, 20.0) // 20% of 1000
	expected := decimal.NewFromFloat(200.00)

	assert.True(t, result.Equal(expected))
}

func TestDivideSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		b        decimal.Decimal
		expected decimal.Decimal
	}{
		{"Normal division", decimal.NewFromFloat(10.00), decimal.NewFromFloat(2.00), decimal.NewFromFloat(5.00)},
		{"Division by zero", decimal.NewFromFloat(10.00), decimal.Zero, decimal.Zero},
		{"Fraction result", decimal.NewFromFloat(1.00), decimal.NewFromFloat(3.00), decimal.NewFromFloat(0.33)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DivideSafe(tt.a, tt.b)
			assert.True(t, result.Equal(tt.expected))
		})
	}
}
