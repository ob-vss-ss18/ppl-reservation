package main

import (
	"net/smtp"
	"strconv"
	"os"
	"fmt"
)

func SendMail(recipients []string, msg []byte) {

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, smtpUser, recipients, msg)

	if err != nil {
		fmt.Println(err)
	}

}
