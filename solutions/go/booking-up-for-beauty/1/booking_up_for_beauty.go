package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	dateHour := t.Hour()

	return dateHour >= 12 && dateHour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	dateFormatted := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.",
		t.Weekday().String(), // %s
		t.Month().String(),   // %s
		t.Day(),              // %d
		t.Year(),             // %d
		t.Hour(),             // %d
		t.Minute(),           // %02d
	)

	return dateFormatted
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "1/2/2006"
	currentYear := fmt.Sprintf("%d", time.Now().Year())
	t, _ := time.Parse(layout, "09/15/"+currentYear)

	return t
}
