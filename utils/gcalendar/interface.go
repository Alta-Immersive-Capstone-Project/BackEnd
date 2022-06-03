package gcalendar

import (
	"kost/entities"

	calendar "google.golang.org/api/calendar/v3"
)

type CalendarControl interface {
	Login(state string) string
	CreateReminder(code string, data entities.DataReminder) (calendar.Event, error)
}
