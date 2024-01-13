package booking

import (
	"fmt"
	"time"
)

func parseDateWithLayouts(date string, layouts ...string) (time.Time, error) {
	var t time.Time
	var err error

	for _, layout := range layouts {
		t, err = time.Parse(layout, date)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, err
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{"1/2/2006 15:04:05", "January 2, 2006 15:04:05", "Monday, January 2, 2006 15:04:05"}
	d, err := parseDateWithLayouts(date, layouts...)
	if err != nil {
		return time.Time{}
	}

	t := time.Date(
		d.Year(), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(), d.Nanosecond(),
		time.UTC,
	)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now().UTC().Truncate(24 * time.Hour)
	return Schedule(date).Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	d := Schedule(date).UTC().Hour()
	return d >= 12 && d <= 17
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	d := Schedule(date).UTC().Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", d)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule("09/15/2024 00:00:00")
}
