package cmd

import "github.com/konoui/go-alfred"

func Initialize(*alfred.Workflow) error {
	if alfred.IsAutoUpdateWorkflowEnabled() {

	}
	return nil
}

func Keyword() string {
	return ""
}
