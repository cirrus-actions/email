package internal

import "os"

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func NeutralExit() {
	// see https://developer.github.com/actions/creating-github-actions/accessing-the-runtime-environment/#exit-codes-and-statuses
	os.Exit(78)
}
