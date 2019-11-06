package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "ratchapong.b5917501@gmail.com", "e575g73wk")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"ratchapong5924@gmail.com"}
	msg := strings.NewReader("To: " + to[0] + "\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "ratchapong.b5917501@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
