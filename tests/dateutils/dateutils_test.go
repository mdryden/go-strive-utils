package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/strivesolutions/go-strive-utils/pkg/dateutils"
	"github.com/strivesolutions/go-strive-utils/tests/testingutils"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxEndDate(t *testing.T) {
	result := dateutils.GetMaxEndDate()
	expectedResult := time.Date(9998, time.December, 31, 23, 59, 59, 0, time.UTC)

	if result != expectedResult {
		t.Fail()
	}
}

func TestFormatSqlDateTime(t *testing.T) {
	maxEndDate := dateutils.GetMaxEndDate()
	result := dateutils.FormatSqlDateTime(maxEndDate)
	expectedResult := "9998-12-31 23:59:59"
	if result != expectedResult {
		t.Fail()
	}
}

var ageTests = []struct {
	birthDate   time.Time
	asOfDate    time.Time
	expectedAge int
}{
	{time.Date(1982, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 8, 0, 0, 0, 0, time.UTC), 40},  // already had birthday this year
	{time.Date(1982, 9, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 8, 0, 0, 0, 0, time.UTC), 39},  // birthday coming up
	{time.Date(1982, 3, 10, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 8, 0, 0, 0, 0, time.UTC), 39}, // birthday coming up this month
	{time.Date(1982, 3, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 8, 0, 0, 0, 0, time.UTC), 40},  // birthday earlier this month
	{time.Date(1982, 3, 8, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 8, 0, 0, 0, 0, time.UTC), 40},  // birthday on as of date
}

func TestGetAgeAsOf(t *testing.T) {
	for i, input := range ageTests {
		expected := input.expectedAge
		actual := dateutils.GetAgeAsOf(input.birthDate, input.asOfDate)
		assert.Equal(t, expected, actual, fmt.Sprintf("Test case %d", i+1))
	}
}

var startOfNextMonthTests = []struct {
	fromDate     time.Time
	inclusive    bool
	expectedDate time.Time
}{
	{testingutils.CreateDate(2022, 1, 1), true, testingutils.CreateDate(2022, 1, 1)},
	{testingutils.CreateDate(2022, 1, 15), true, testingutils.CreateDate(2022, 2, 1)},
	{testingutils.CreateDate(2022, 1, 1), false, testingutils.CreateDate(2022, 2, 1)},
	{testingutils.CreateDate(2022, 1, 15), false, testingutils.CreateDate(2022, 2, 1)},
	{testingutils.CreateDate(2022, 1, 31), true, testingutils.CreateDate(2022, 2, 1)},
	{testingutils.CreateDate(2022, 1, 31), false, testingutils.CreateDate(2022, 2, 1)},
	{testingutils.CreateDate(2022, 2, 1), true, testingutils.CreateDate(2022, 2, 1)}, // Feb being short deserves it's own test cases
	{testingutils.CreateDate(2022, 2, 1), false, testingutils.CreateDate(2022, 3, 1)},
	{testingutils.CreateDate(2022, 2, 16), true, testingutils.CreateDate(2022, 3, 1)},
	{testingutils.CreateDate(2022, 2, 13), false, testingutils.CreateDate(2022, 3, 1)},
	{testingutils.CreateDate(2022, 12, 1), true, testingutils.CreateDate(2022, 12, 1)},  // Make sure the year doesn't roll over
	{testingutils.CreateDate(2022, 12, 1), false, testingutils.CreateDate(2023, 1, 1)},  // Make sure the year rolls over properly.
	{testingutils.CreateDate(2022, 12, 16), true, testingutils.CreateDate(2023, 1, 1)},  // Make sure the year rolls over properly.
	{testingutils.CreateDate(2022, 12, 13), false, testingutils.CreateDate(2023, 1, 1)}, // Make sure the year rolls over properly.
}

func TestStartOfNextMonth(t *testing.T) {
	for i, input := range startOfNextMonthTests {
		expected := input.expectedDate
		actual := dateutils.StartOfNextMonth(input.fromDate, input.inclusive)

		assert.Equal(t, expected, actual, fmt.Sprintf("Test case %d", i+1))
	}
}

var birthDayForAgeTests = []struct {
	birthDate    time.Time
	targetAge    int
	expectedDate time.Time
}{
	{testingutils.CreateDate(1980, 1, 1), 40, testingutils.CreateDate(2020, 1, 1)},
	{testingutils.CreateDate(1995, 6, 15), 20, testingutils.CreateDate(2015, 6, 15)},
	{testingutils.CreateDate(1984, 1, 1), 50, testingutils.CreateDate(2034, 1, 1)},
	{testingutils.CreateDate(1984, 2, 29), 50, testingutils.CreateDate(2034, 3, 1)},  // born on feb 29th, non-leap year birthday
	{testingutils.CreateDate(1984, 2, 29), 48, testingutils.CreateDate(2032, 2, 29)}, // born on feb 29th, leap year birthday
	{testingutils.CreateDate(1982, 8, 10), 50, testingutils.CreateDate(2032, 8, 10)},
}

func TestGetBirthdayForAge(t *testing.T) {
	for i, input := range birthDayForAgeTests {
		expected := input.expectedDate
		actual := dateutils.GetBirthdayForAge(input.birthDate, input.targetAge)

		assert.Equal(t, expected, actual, fmt.Sprintf("Test case %d", i+1))
	}
}

var decimalAgeTests = []struct {
	birthDate   time.Time
	asOfDate    time.Time
	expectedAge float64
}{
	{testingutils.CreateDate(1982, 8, 10), testingutils.CreateDate(2005, 12, 31), 23.3922},
	{testingutils.CreateDate(1982, 8, 10), testingutils.CreateDate(2032, 9, 1), 50.0616},
	{testingutils.CreateDate(1961, 9, 16), testingutils.CreateDate(1990, 12, 31), 29.2895},
	{testingutils.CreateDate(1961, 9, 16), testingutils.CreateDate(2021, 11, 1), 60.1259},
	{testingutils.CreateDate(1959, 7, 10), testingutils.CreateDate(2012, 12, 31), 53.4784},
	{testingutils.CreateDate(1982, 8, 10), testingutils.CreateDate(2022, 3, 18), 39.6030},
	{testingutils.CreateDate(1985, 8, 10), testingutils.CreateDate(2022, 3, 9), 36.5777},
	{testingutils.CreateDate(1993, 5, 1), testingutils.CreateDate(2043, 5, 1), 50.0000},
}

func TestDecimalAgeAsOf(t *testing.T) {
	for i, input := range decimalAgeTests {
		expected := input.expectedAge
		actual := dateutils.DecimalAgeAsOf(input.birthDate, input.asOfDate)

		assert.Equal(t, expected, actual, fmt.Sprintf("Test case %d", i))
	}
}
