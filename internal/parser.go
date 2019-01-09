package internal

import (
	"encoding/json"
	"github.com/google/go-github/v21/github"
	"io/ioutil"
)

func Parse(path string) (github.CheckSuiteEvent, error) {
	event := github.CheckSuiteEvent{}
	eventBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return event, err
	}

	err = json.Unmarshal(eventBytes, &event)
	return event, err
}
