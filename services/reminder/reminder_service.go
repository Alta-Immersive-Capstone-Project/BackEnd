package reminder

import (
	"errors"
	"kost/repositories/transactions"
	"kost/utils/gcalendar"

	"github.com/labstack/gommon/log"
	"google.golang.org/api/calendar/v3"
)

type ServiceReminder struct {
	cal   gcalendar.CalendarControl
	trans transactions.TransactionModel
}

func NewReminderServices(gcalender gcalendar.CalendarControl, transModel transactions.TransactionModel) *ServiceReminder {
	return &ServiceReminder{
		cal:   gcalender,
		trans: transModel,
	}
}

func (sr *ServiceReminder) GetLoginUrl(state string) string {
	authUrl := sr.cal.Login(state)
	return authUrl
}

func (sr *ServiceReminder) CreateEvent(code string, BookingID string) (calendar.Event, error) {

	data, err := sr.trans.GetTransactionByBookingID(BookingID)
	if err != nil {
		return calendar.Event{}, errors.New("Data Not Found")
	}
	result, err := sr.cal.CreateReminder(code, data)
	if err != nil {
		log.Warn(err)
		return calendar.Event{}, errors.New("Error Create Reminder")
	}

	return result, nil
}
