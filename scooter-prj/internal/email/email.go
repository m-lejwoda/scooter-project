package email

import (
	"fmt"
	"net/smtp"
)

type Mailer interface {
	SendResetPassword(toEmail, rawToken string) error
}

type SMTPMailer struct {
	addr string
}

func NewSMTPMailer(addr string) *SMTPMailer {
	return &SMTPMailer{addr: addr}
}

func (m *SMTPMailer) SendResetPassword(toEmail, rawToken string) error {
	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", rawToken)

	msg := []byte(fmt.Sprintf(
		"From: no-reply@scooter.local\r\n"+
			"To: %s\r\n"+
			"Subject: Reset hasla\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
			"Kliknij link, aby zresetować hasło: <a href=\"%s\">Resetuj hasło</a>",
		toEmail, resetLink,
	))

	return smtp.SendMail(m.addr, nil, "no-reply@scooter.local", []string{toEmail}, msg)
}
