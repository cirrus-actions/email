package main

import (
	"github.com/cirrus-actions/email/internal"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	var spec internal.Specification
	err := envconfig.Process("mail", &spec)
	if err != nil {
		log.Fatal(err.Error())
	}
	if spec.MailPort == 0 {
		spec.MailPort = 587
	}
	internal.SendNotification(spec)
}
