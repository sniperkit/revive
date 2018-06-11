package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/rule"
)

// TestSuperfluousElse rule.
func TestSuperfluousElse(t *testing.T) {
	testRule(t, "superfluous-else", &rule.SuperfluousElseRule{})
}
