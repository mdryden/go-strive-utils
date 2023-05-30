package mathutils

import (
	"math"

	"github.com/shopspring/decimal"
)

func Prorate(amount decimal.Decimal, numerator float64, denominator float64) decimal.Decimal {
	percentage := decimal.NewFromFloat(numerator).Div(decimal.NewFromFloat(denominator))
	return amount.Mul(percentage)
}

func RoundFloat2(value float64) float64 {
	return math.Round(value*100) / 100
}

func RoundFloat4(value float64) float64 {
	return math.Round(value*10000) / 10000
}

func RoundUpToInt(value float64) int {
	rounded := math.Round(value)
	if rounded < value {
		rounded++
	}
	return int(rounded)
}
