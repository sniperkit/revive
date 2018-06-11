package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/lint"
	"github.com/sniperkit/revive/pkg/rule"
)

func TestArgumentLimit(t *testing.T) {
	testRule(t, "argument-limit", &rule.ArgumentsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
