package tz

import (
	"github.com/dromara/carbon/v2"
)

// UTCDateToTz converts a UTC date string to the specified timezone.
// It accepts date strings in "YYYY-MM-DD" format.
//
// Example:
//
//	dateInNewYork, err := UTCDateToTz("2022-01-01", "America/New_York")
//
// Parameters:
//   - utcDateString: a string representing the UTC date in "YYYY-MM-DD" format.
//   - timezone: a string representing the timezone (e.g., "America/New_York").
//
// Returns:
//   - a string representing the converted date in "YYYY-MM-DD" format, or an empty string and an error if parsing fails.
func UTCDateToTz(utcDateString string, timezone string) (string, error) {
	parsedDate := carbon.Parse(utcDateString, carbon.UTC)
	dateInTz := parsedDate.SetTimezone(timezone)
	return dateInTz.ToDateString(), nil
}
