package sendto

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Mail struct {
	From    EmailAddress
	To      []string `json:"to,omitempty"`
	Subject string   `json:"subject,omitempty"`
	Body    string   `json:"body,omitempty"`
}

func BuildMessage(mail Mail) string {
	msg := "IME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg += fmt.Sprintf("From: %s<%s>\r\n", mail.From.Name, mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	msg += fmt.Sprintf("\r\n")

	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprint("Your OTP is %s, Please enter it to verify your account.", otp),
	}

	// message := BuildMessage(contentEmail)
	//
	// auth := smtp.PlainAuth
	return nil
}
