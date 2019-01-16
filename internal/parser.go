package internal

import (
	"encoding/json"
	"github.com/google/go-github/v21/github"
	"io/ioutil"
)

type checkSuiteEventExt struct {
	CheckSuite *CheckSuiteExt `json:"check_suite,omitempty"`
}

type CheckSuiteExt struct {
	HeadCommit *github.Commit `json:"head_commit,omitempty"`
}

func Parse(path string) (github.CheckSuiteEvent, github.Commit, error) {
	event := github.CheckSuiteEvent{}
	eventExt := checkSuiteEventExt{}
	eventBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return event, *eventExt.CheckSuite.HeadCommit, err
	}

	err = json.Unmarshal(eventBytes, &event)

	if err != nil {
		return event, *eventExt.CheckSuite.HeadCommit, err
	}

	err = json.Unmarshal(eventBytes, &eventExt)
	return event, *eventExt.CheckSuite.HeadCommit, err
}
