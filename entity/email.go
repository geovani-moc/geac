package entity

import (
	"log"
	"net/smtp"
)

// Email struct
type Email struct {
	from       string
	password   string
	smtpHost   string
	smtpPort   string
	Attachment Certificate
	To         []string
	Subject    string
	Body       []byte
}

// Send function to invite email
func (email *Email) Send() {
	auth := smtp.PlainAuth("", email.from, email.password, email.smtpHost)

	err := smtp.SendMail(email.smtpHost+":"+email.smtpPort, auth, email.from, email.To, email.Body)
	if nil != err {
		log.Println("Error to invite email", err)
		return
	}
	log.Println("Email sent succesfully.")
}
