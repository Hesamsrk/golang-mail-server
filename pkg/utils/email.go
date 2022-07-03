package utils

import (
	"fmt"
	"net/smtp"

	"github.com/Hesamsrk/golang-mail-server/pkg/config"
)

func SendMail(to string, message string) error {

	// Sender data.
	from := config.LocalConfig.EMAIL_ADDRESS
	password := config.LocalConfig.EMAIL_PASSWORD

	// smtp server configuration.
	smtpHost := "smtp.zoho.com"
	smtpPort := "587"

	// Message.
	Message := []byte(message)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, Message)
	if err != nil {
		return err
	}
	fmt.Println("******************************************************")
	fmt.Println("Email Sent:")
	fmt.Println("From: " + from)
	fmt.Println("To: " + to)
	fmt.Println("Message:")
	fmt.Println(message)
	fmt.Println("******************************************************")
	return nil
}
