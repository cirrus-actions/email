package internal

import (
	"testing"
)

func Test_Parsing(t *testing.T) {
	event, commit, err := Parse("../testdata/check_suite_completed.json")
	if err != nil {
		t.Error(err)
	}
	checkSuite := event.CheckSuite
	repository := event.Repo
	if repository.Owner.Login == nil {
		t.Errorf("Nil!")
	}
	if repository.Name == nil {
		t.Errorf("Nil!")
	}
	if checkSuite.HeadSHA == nil {
		t.Errorf("Nil!")
	}
	if commit.Author == nil {
		t.Errorf("Nil!")
	}
}
