package internal

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

type Specification struct {
	EventPath    string `envconfig:"GITHUB_EVENT_PATH",required:"true"`
	MailHost     string `envconfig:"MAIL_HOST",required:"true"`
	MailPort     int    `envconfig:"MAIL_PORT"`
	MailFrom     string `envconfig:"MAIL_FROM",required:"true"`
	MailUsername string `envconfig:"MAIL_USERNAME",required:"true"`
	MailPassword string `envconfig:"MAIL_PASSWORD",required:"true"`
}

func SendNotification(spec Specification) {
	event, commit, err := Parse(spec.EventPath)
	if err != nil {
		log.Fatalf("Failed to parse event! %s", err)
	}

	fmt.Printf("Dialing %v:%v...\n", spec.MailHost, spec.MailPort)
	var dialer *gomail.Dialer
	if spec.MailUsername == "" && spec.MailPassword == "" {
		dialer = &gomail.Dialer{Host: spec.MailHost, Port: spec.MailPort}
	} else {
		dialer = gomail.NewDialer(spec.MailHost, spec.MailPort, spec.MailUsername, spec.MailPassword)
	}
	message, err := generateEmail(spec, event, commit)
	if err != nil {
		log.Fatalf("Failed to generate email! %s", err)
	}
	err = dialer.DialAndSend(message)
	if err != nil {
		log.Fatalf("Failed to send email! %s", err)
	}
}
