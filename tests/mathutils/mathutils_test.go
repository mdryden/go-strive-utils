package tests

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-strive-utils/pkg/mathutils"
	"github.com/strivesolutions/go-strive-utils/tests/testingutils"
)

func TestProrate25Pct(t *testing.T) {
	amount := decimal.NewFromFloat32(100.0)

	expected := decimal.NewFromFloat(25.0)
	actual := mathutils.Prorate(amount, 2.5, 10)

	testingutils.AssertDecimalsEqual(t, expected, actual)
}

func TestProrate50Pct(t *testing.T) {
	amount := decimal.NewFromFloat32(100.0)

	expected := decimal.NewFromFloat(50.0)
	actual := mathutils.Prorate(amount, 5, 10)

	testingutils.AssertDecimalsEqual(t, expected, actual)
}

func TestProrate75Pct(t *testing.T) {
	amount := decimal.NewFromFloat32(100.0)

	expected := decimal.NewFromFloat(75.0)
	actual := mathutils.Prorate(amount, 7.5, 10)

	testingutils.AssertDecimalsEqual(t, expected, actual)
}

var roundUpIntCases = []struct {
	value    float64
	expected int
}{
	{1.0, 1},
	{1.1, 2},
	{1.2, 2},
	{1.3, 2},
	{1.4, 2},
	{1.5, 2},
	{1.6, 2},
	{1.7, 2},
	{1.8, 2},
	{1.9, 2},
	{2.0, 2},
	{2.1, 3},
	{2.2, 3},
	{2.3, 3},
	{2.4, 3},
	{2.5, 3},
	{2.6, 3},
	{2.7, 3},
	{2.8, 3},
	{2.9, 3},
}

func TestRoundUpInt(t *testing.T) {
	for _, input := range roundUpIntCases {
		expected := input.expected
		actual := mathutils.RoundUpToInt(input.value)

		assert.Equal(t, expected, actual)
	}
}
