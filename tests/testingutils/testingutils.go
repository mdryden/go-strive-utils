package testingutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func AssertDecimalsEqual(t *testing.T, expected decimal.Decimal, actual decimal.Decimal, msgAndArgs ...interface{}) bool {
	if !expected.Equals(actual) {
		return assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected.String(), actual.String()), msgAndArgs...)

	}

	return true
}

func CreateDate(year int, month int, day int) time.Time {
	monthAsMonth := time.Month(month)
	return time.Date(year, monthAsMonth, day, 0, 0, 0, 0, time.UTC)
}
