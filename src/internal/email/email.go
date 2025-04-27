package email

import (
	"gopkg.in/gomail.v2" // Import the gomail library
)

type Email struct {
	from     string
	to       string
	password string
	subject  string
}

// inputs are:
// from: email address
// password: app password
func NewEmail(from, to, subject, password string) *Email {
	return &Email{
		from:     from,
		to:       to,
		password: password,
		subject:  subject,
	}
}

func (e *Email) Send(htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", e.to)
	m.SetHeader("Subject", e.subject)

	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, e.from, e.password)

	return d.DialAndSend(m)
}
