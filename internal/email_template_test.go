package internal

import (
	"testing"
)

func Test_Subject(t *testing.T) {
	event, _, err := Parse("../testdata/check_suite_completed.json")
	if err != nil {
		t.Error(err)
	}
	subject, err := Render(DefaultSubjectTemplate, event)
	if err != nil {
		t.Errorf("could not render subject template: %v", err)
	}
	if subject != "Cirrus CI check for cirrus-actions/email#master completed" {
		t.Errorf("Wrong subject: %s", subject)
	}
}
