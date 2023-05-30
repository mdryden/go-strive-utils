package dateutils

import (
	"time"

	"github.com/strivesolutions/go-strive-utils/pkg/mathutils"
)

func GetMaxEndDate() time.Time {
	return time.Date(9998, time.December, 31, 23, 59, 59, 0, time.UTC)
}

func FormatSqlDateTime(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}

func ToEndOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 0, time.UTC)
}

func ToStartOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.UTC)
}

// GetCurrentAge calculates age in years for a given birthdate, as of right now.
// Note, this function is not timezone aware
func GetCurrentAge(birthDate time.Time) int {
	return GetAgeAsOf(birthDate, time.Now())
}

// GetCurrentAge calculates age in years for a given birthdate, as of a specific date.
// Note, this function is not timezone aware
func GetAgeAsOf(birthDate time.Time, asOfDate time.Time) int {
	if birthDate.After(asOfDate) {
		tmp := asOfDate
		asOfDate = birthDate
		birthDate = tmp
	}

	years := asOfDate.Year() - birthDate.Year()
	months := asOfDate.Month() - birthDate.Month()

	birthDateLaterThisYear := months < 0
	if birthDateLaterThisYear {
		years--
	}

	birthDateThisMonth := months == 0
	if birthDateThisMonth {
		days := asOfDate.Day() - birthDate.Day()
		birthDateLaterThisMonth := days < 0

		if birthDateLaterThisMonth {
			years--
		}
	}

	return years
}

func DecimalAgeAsOf(birthDate time.Time, asOfDate time.Time) float64 {
	var years float64
	if birthDate.Day() == asOfDate.Day() && birthDate.Month() == asOfDate.Month() {
		years = float64(asOfDate.Year() - birthDate.Year())
	} else {
		days := asOfDate.Sub(birthDate).Hours() / 24.0
		years = days / 365.25
	}
	return mathutils.RoundFloat4(years)
}

func ParseIsoString(input string) (time.Time, error) {
	// This is here so we can support more formats (eg: microseconds) in the future
	return time.Parse("2006-01-02T15:04:05Z", input)
}

// Returns the first of the next month, after the provided fromDate, ignoring the time.
// If inclusive == true and the fromDate is the first of the month, it will return the fromDate.
// If inclusive == false and the fromDate is the first of the month, it will return fromDate + 1 month.
func StartOfNextMonth(fromDate time.Time, inclusive bool) time.Time {
	fromDate = ToStartOfDay(fromDate)

	dayIsFirst := fromDate.Day() == 1

	if dayIsFirst && inclusive {
		return fromDate
	}

	nextYear := fromDate.Year()

	if fromDate.Month() == time.December {
		nextYear++
	}

	nextMonth := fromDate.Month() + 1
	if int(nextMonth) > 12 {
		nextMonth = time.January
	}

	startOfNextMonth := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, fromDate.Location())

	return startOfNextMonth
}

func GetBirthdayForAge(birthDate time.Time, targetAge int) time.Time {
	return birthDate.AddDate(targetAge, 0, 0)
}

func GetDateForAge(birthDate time.Time, targetAge float64) time.Time {
	years := int(targetAge)
	days := int(365.25 * (targetAge - float64(years)))

	return birthDate.AddDate(years, 0, days)
}
