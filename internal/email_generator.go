package internal

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/google/go-github/v21/github"
	"gopkg.in/gomail.v2"
)

func generateEmail(spec Specification, event github.CheckSuiteEvent, commit github.Commit) (*gomail.Message, error) {
	message := gomail.NewMessage()

	message.SetAddressHeader("From", spec.MailFrom, *event.CheckSuite.App.Name)
	message.SetAddressHeader("To", *commit.Committer.Email, *commit.Committer.Name)
	fmt.Printf("Creating email for %v <%v>...\n", *commit.Committer.Name, *commit.Committer.Email)

	// populate subject:
	subject, err := Render(DefaultSubjectTemplate, event)
	if err != nil {
		return nil, fmt.Errorf("could not render subject template: %v", err)
	}
	message.SetHeader("Subject", subject)

	contentMD, err := Render(DefauleEmailMarkdownTemplate, event)
	if err != nil {
		return nil, fmt.Errorf("could not render content template: %v", err)
	}
	message.AddAlternative("text/text", contentMD)

	contentHTML := markdown.ToHTML([]byte(contentMD), nil, nil)
	message.AddAlternative("text/html", string(contentHTML))
	return message, nil
}
