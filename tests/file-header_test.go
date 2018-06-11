package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/lint"
	"github.com/sniperkit/revive/pkg/rule"
)

func TestLintFileHeader(t *testing.T) {
	testRule(t, "lint-file-header1", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header2", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header3", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header4", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"^\\sfoobar$"},
	})

	testRule(t, "lint-file-header5", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"^\\sfoo.*bar$"},
	})
}
