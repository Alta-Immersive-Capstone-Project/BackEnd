package email

type EmailControl interface {
	SendEmail(sender string, subject string, body string, recipient string) (string, error)
}
