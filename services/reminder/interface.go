package reminder

import (
	"google.golang.org/api/calendar/v3"
)

type ReminderService interface {
	CreateEvent(code string, BookingID string) (calendar.Event, error)
	GetLoginUrl(state string) string
}
