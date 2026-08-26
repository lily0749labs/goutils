package goutils_test

import (
	"testing"

	"github.com/lily0749labs/goutils"
)

func TestNewFacades(t *testing.T) {
	t.Parallel()

	if got := goutils.Money.FormatCents(-1); got != "-0.01" {
		t.Fatalf("Money.FormatCents() = %q, want -0.01", got)
	}
	if got := goutils.String.Mask("13800138000", 3, 4); got != "138****8000" {
		t.Fatalf("String.Mask() = %q, want 138****8000", got)
	}
}
