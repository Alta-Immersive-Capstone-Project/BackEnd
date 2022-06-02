package reminder

import "google.golang.org/api/calendar/v3"

type ReminderService interface {
	GetListEvent() (*calendar.Event, error)
}
