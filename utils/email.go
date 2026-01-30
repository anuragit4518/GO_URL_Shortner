package utils

import (
	"fmt"
	"net/smtp"
)

// SendEmail sends an email using SMTP
func SendEmail(toEmail, subject, body string) error {
	from := "anurag120307@gmail.com"
	password := "qcjw snjy haqu vezd"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\n\n%s",
		from,
		toEmail,
		subject,
		body,
	)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{toEmail},
		[]byte(message),
	)

	return err
}
