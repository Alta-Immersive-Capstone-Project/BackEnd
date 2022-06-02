package email

import (
	"kost/configs"

	"github.com/mailgun/mailgun-go"
)

type EmailService struct {
	email mailgun.MailgunImpl
}

func NewEmailConfig() *EmailService {

	mg := mailgun.NewMailgun(configs.Get().Email.Domain, configs.Get().Email.ApiKey)
	return &EmailService{
		email: *mg,
	}
}

func (emailService EmailService) SendEmail(sender string, subject string, body string, recipient string) (string, error) {

	// The message object allows you to add attachments and Bcc recipients
	message := emailService.email.NewMessage(sender, subject, body, recipient)

	// Send the message with a 10 second timeout
	_, id, err := emailService.email.Send(message)

	if err != nil {
		return "", err
	}

	return id, nil
}
