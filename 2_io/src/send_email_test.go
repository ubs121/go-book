package io

import (
	"bytes"
	"log"
	"net/smtp"
	"testing"
)

func TestSMTP(t *testing.T) {
	// SMTP сервер рүү холбогдох
	client, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}
	// илгээгч, хүлээн авагчийг тохируулах
	client.Mail("sender@example.org")
	client.Rcpt("recipient@example.net")

	// э-мэйлийн бие хэсэг үүсгэх
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("Э-мэйлийн бие.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}

func TestSMTPAuth(t *testing.T) {
	// нэвтрэх эрхийг тохируулах
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)

	// э-мэйл илгээх
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("Э-мэйлийн бие."),
	)

	if err != nil {
		log.Fatal(err)
	}
}
