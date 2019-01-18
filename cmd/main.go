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
		log.Fatalf("Failed to get specification from environment variables: %v", err)
	}
	internal.SendNotification(spec)
}
