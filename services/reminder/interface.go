package reminder

import (
	"kost/entities"

	"google.golang.org/api/calendar/v3"
)

type ReminderService interface {
	CreateEvent(reminder *entities.Reminder) *calendar.Event
	// GetListEvent() (*calendar.Event, error)
	GetLoginUrl(state string) string
	GetUserDataFromGoogle(code string) ([]byte, error)
}
