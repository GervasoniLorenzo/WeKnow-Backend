package adapter

import (
	"fmt"
	"net/smtp"
	"weKnow/model"
)

type MailerAdapter interface {
	SendEmail(emailData model.Email) error
}

func (a *KnownAdapter) SendEmail(emailData model.Email) error {

	from := "your-email@example.com"
	message := []byte("Subject: " + emailData.Subject + "\r\n" +
		"To: " + emailData.To + "\r\n" +
		"From: " + from + "\r\n" +
		"\r\n" + // Blank line separating headers from body
		emailData.Body)

	// Connect to the SMTP server
	// auth := smtp.PlainAuth("", from, a.config.Email.Password, a.config.Email.SmtpHost)
	err := smtp.SendMail(a.config.Email.SmtpHost+":"+a.config.Email.SmtpPort, nil, from, []string{emailData.To}, message)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil

}
