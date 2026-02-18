package mail

import (
	"fmt"
	"net/smtp"
)

type UserMail interface {
	SendEmailToUser(to []string, subject string, body string) error
}

type SmtpMail struct {
	addr string
	auth smtp.Auth
}

func NewSmtpMailer(host string, port int, user string, pass string) UserMail {
	return &SmtpMail{
		addr: fmt.Sprintf("%s:%d", host, port),
		auth: smtp.PlainAuth("", user, pass, host),
	}
}

func (s *SmtpMail) SendEmailToUser(to []string, subject string, body string) error {
	msg := []byte("Subject: " + subject + "\r\n\r\n" + body)
	return smtp.SendMail(s.addr, s.auth, "sender@example.com", to, msg)
}
