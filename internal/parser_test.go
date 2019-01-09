package internal

import (
	"testing"
)

func Test_Parsing(t *testing.T) {
	_, err := Parse("../testdata/check_suite_completed.json")
	if err != nil {
		t.Error(err)
	}
}
