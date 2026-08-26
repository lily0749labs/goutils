package money

import (
	"math"
	"testing"
)

func TestFormatCents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		cents int64
		want  string
	}{
		{name: "zero", cents: 0, want: "0.00"},
		{name: "one cent", cents: 1, want: "0.01"},
		{name: "whole amount", cents: 100, want: "1.00"},
		{name: "positive", cents: 12345, want: "123.45"},
		{name: "negative cent", cents: -1, want: "-0.01"},
		{name: "negative", cents: -12345, want: "-123.45"},
		{name: "minimum int64", cents: math.MinInt64, want: "-92233720368547758.08"},
		{name: "maximum int64", cents: math.MaxInt64, want: "92233720368547758.07"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatCents(tc.cents); got != tc.want {
				t.Fatalf("FormatCents(%d) = %q, want %q", tc.cents, got, tc.want)
			}
		})
	}
}
