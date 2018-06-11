package test

import (
	"testing"

	"github.com/sniperkit/revive/pkg/rule"
)

// TestEmptyBlock rule.
func TestEmptyBlock(t *testing.T) {
	testRule(t, "empty-block", &rule.EmptyBlockRule{})
}
