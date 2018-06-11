package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/lint"
	"github.com/sniperkit/revive/pkg/rule"
)

func TestCyclomatic(t *testing.T) {
	testRule(t, "cyclomatic", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
	testRule(t, "cyclomatic-2", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
