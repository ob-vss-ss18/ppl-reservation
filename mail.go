package main

import (
	"net/smtp"
	"strconv"
	"os"
	"log"
)

func SendMail(recipients []string, msg []byte) {

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")

	if err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, smtpUser, recipients, msg)

	if err != nil {
		log.Fatal(err)
	}

}
