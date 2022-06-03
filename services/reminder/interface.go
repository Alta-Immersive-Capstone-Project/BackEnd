package reminder

import (
	"kost/entities"
	"net/http"

	"google.golang.org/api/calendar/v3"
)

type ReminderService interface {
	CreateEvent(code string, sreminder entities.Reminder) (*calendar.Event, error)
	// GetListEvent() (*calendar.Event, error)
	GetLoginUrl(state string) string
	GetUserDataFromGoogle(w http.ResponseWriter, code string) ([]byte, error)
}
