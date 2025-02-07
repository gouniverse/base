package tz

import (
	"github.com/dromara/carbon/v2"
)

// UTCTimeToTz converts a UTC time string to the specified timezone.
// It accepts time strings in "HH:MM" format.
//
// Example:
//
//	timeInNewYork, err := UTCTimeToTz("10:00", "America/New_York")
//
// Parameters:
//   - utcTimeString: a string representing the UTC time in "HH:MM" format.
//   - timezone: a string representing the timezone (e.g., "America/New_York").
//
// Returns:
//   - a string representing the converted time in "HH:MM" format, or an empty string and an error if parsing fails.
func UTCTimeToTz(utcTimeString string, timezone string) string {
	datetime := "2000-01-01 " + utcTimeString // the year is irrelevant, as we only return the time
	parsedTime := carbon.Parse(datetime, carbon.UTC)
	return parsedTime.SetTimezone(timezone).Format("H:i")
}
