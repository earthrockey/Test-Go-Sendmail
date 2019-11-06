package main

import (
	"log"
	"net/smtp"
)

func main() {
	Sendmail("ratchapong.b5917501@gmail.com", "Test SMTP Client", "This is the email body.")
}

func Sendmail(to string, subject string, body string) {
	from := "ratchapong.b5917501@gmail.com"
	pass := "e575g73wk"
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n\n" + body
	if err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg)); err != nil {
		log.Printf("smtp error: %s", err)
	}
}
