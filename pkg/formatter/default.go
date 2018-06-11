package formatter

import (
	"fmt"

	lint "github.com/sniperkit/revive/pkg/lint"
)

// Default is an implementation of the Formatter interface
// which formats the errors to JSON.
type Default struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter
func (f *Default) Name() string {
	return "default"
}

// Format formats the failures gotten from the lint.
func (f *Default) Format(failures <-chan lint.Failure, config lint.RulesConfig) (string, error) {
	for failure := range failures {
		fmt.Printf("%v: %s\n", failure.Position.Start, failure.Failure)
	}
	return "", nil
}
