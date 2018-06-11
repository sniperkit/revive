package formatter

import (
	lint "github.com/sniperkit/revive/pkg/lint"
)

func severity(config lint.RulesConfig, failure lint.Failure) lint.Severity {
	if config, ok := config[failure.RuleName]; ok && config.Severity == lint.SeverityError {
		return lint.SeverityError
	}
	return lint.SeverityWarning
}
