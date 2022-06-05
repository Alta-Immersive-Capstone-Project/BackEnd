package reminder

import (
	"errors"
	"kost/entities"
	mocksTrans "kost/mocks/repositories/transactions"
	mocks "kost/mocks/utils/gcalendar"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/calendar/v3"
)

func TestGetLoginUrl(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		CalendarControl := mocks.NewCalendarControl(t)
		Trans := mocksTrans.NewTransactionModel(t)
		CalendarControl.On("Login", "State").Return("url").Once()

		calenderService := NewReminderServices(CalendarControl, Trans)

		result := calenderService.GetLoginUrl("State")
		assert.Equal(t, "url", result)

	})
}

func TestCreateEvent(t *testing.T) {
	data := entities.DataReminder{}
	booking := "123"
	code := "trial"
	Respond := calendar.Event{Description: "Trial"}
	t.Run("Success Create Event", func(t *testing.T) {
		CalendarControl := mocks.NewCalendarControl(t)
		Trans := mocksTrans.NewTransactionModel(t)
		CalendarControl.On("CreateReminder", code, mock.Anything).Return(Respond, nil)
		Trans.On("GetTransactionByBookingID", booking).Return(data, nil)
		calenderService := NewReminderServices(CalendarControl, Trans)

		result, err := calenderService.CreateEvent(code, booking)
		assert.Equal(t, Respond, result)
		assert.NoError(t, err)
	})
	t.Run("Error Get Trans", func(t *testing.T) {
		CalendarControl := mocks.NewCalendarControl(t)
		Trans := mocksTrans.NewTransactionModel(t)
		Trans.On("GetTransactionByBookingID", "123").Return(data, errors.New("Error")).Once()
		calenderService := NewReminderServices(CalendarControl, Trans)

		result, err := calenderService.CreateEvent(code, booking)
		assert.NotEqual(t, Respond, result)
		assert.Error(t, err)
	})
	t.Run("Error Create Event", func(t *testing.T) {
		CalendarControl := mocks.NewCalendarControl(t)
		Trans := mocksTrans.NewTransactionModel(t)
		Trans.On("GetTransactionByBookingID", booking).Return(data, nil)
		CalendarControl.On("CreateReminder", code, mock.Anything).Return(Respond, errors.New("Error"))
		calenderService := NewReminderServices(CalendarControl, Trans)

		result, err := calenderService.CreateEvent(code, booking)
		assert.NotEqual(t, Respond, result)
		assert.Error(t, err)
	})
}
