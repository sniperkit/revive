package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/lint"
	"github.com/sniperkit/revive/pkg/rule"
)

func TestDisabledAnnotations(t *testing.T) {
	testRule(t, "disable-annotations", &rule.ExportedRule{}, &lint.RuleConfig{})
}
