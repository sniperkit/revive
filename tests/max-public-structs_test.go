package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/lint"
	"github.com/sniperkit/revive/pkg/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
