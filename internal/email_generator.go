package internal

import (
	"context"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
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

	githubClient := newGitHubClient(spec)
	listCheckRunsResults, _, err := githubClient.Checks.ListCheckRunsCheckSuite(context.Background(), *event.Repo.Owner.Login, *event.Repo.Name, *event.CheckSuite.ID, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get check runs: %v", err)
	}

	contentBuilder := strings.Builder{}
	for _, checkRun := range listCheckRunsResults.CheckRuns {
		var duraiton time.Duration

		if checkRun.CompletedAt != nil {
			duraiton = checkRun.CompletedAt.Sub(checkRun.StartedAt.Time)
		}

		data := struct {
			CheckRun *github.CheckRun
			Duration time.Duration
		}{
			CheckRun: checkRun,
			Duration: duraiton,
		}
		contentPart, err := Render(DefauleEmailMarkdownTemplate, data)
		if err != nil {
			return nil, fmt.Errorf("could not render content template for check run '%s': %v", *checkRun.Name, err)
		}
		contentBuilder.WriteString(contentPart)
		contentBuilder.WriteString("\n")
	}

	contentMD := contentBuilder.String()
	message.AddAlternative("text/text", contentMD)
	contentHTML := markdown.ToHTML([]byte(contentMD), nil, nil)
	message.AddAlternative("text/html", string(contentHTML))
	return message, nil
}

func newGitHubClient(spec Specification) *github.Client {
	if spec.GitHubToken == "" {
		return github.NewClient(nil)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: spec.GitHubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
