package floatutil

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestTruncate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		value     float64
		precision int
		want      float64
	}{
		{name: "positive", value: 1.239, precision: 2, want: 1.23},
		{name: "negative", value: -1.239, precision: 2, want: -1.23},
		{name: "does not pad", value: 1.2, precision: 2, want: 1.2},
		{name: "does not round", value: 1.9999999999999998, precision: 2, want: 1.99},
		{name: "integer", value: 12.99, precision: 0, want: 12},
		{name: "negative precision", value: -12.99, precision: -1, want: -12},
		{name: "zero", value: 0, precision: 2, want: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := Float.Truncate(testCase.value, testCase.precision); got != testCase.want {
				t.Fatalf("Float.Truncate(%v, %d) = %v, want %v", testCase.value, testCase.precision, got, testCase.want)
			}
		})
	}
}

func TestTruncateSpecialValues(t *testing.T) {
	t.Parallel()

	if got := Float.Truncate(math.Inf(1), 2); !math.IsInf(got, 1) {
		t.Fatalf("Float.Truncate(+Inf, 2) = %v", got)
	}
	if got := Float.Truncate(math.Inf(-1), 2); !math.IsInf(got, -1) {
		t.Fatalf("Float.Truncate(-Inf, 2) = %v", got)
	}
	if got := Float.Truncate(math.NaN(), 2); !math.IsNaN(got) {
		t.Fatalf("Float.Truncate(NaN, 2) = %v", got)
	}
}

func TestTruncateString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		value     float64
		precision int
		want      string
	}{
		{value: 1.239, precision: 2, want: "1.23"},
		{value: -1.239, precision: 2, want: "-1.23"},
		{value: 1.2, precision: 2, want: "1.2"},
		{value: 12.99, precision: 0, want: "12."},
		{value: 12.99, precision: -1, want: "12."},
	}
	for _, testCase := range testCases {
		if got := Float.TruncateString(testCase.value, testCase.precision); got != testCase.want {
			t.Errorf("Float.TruncateString(%v, %d) = %q, want %q", testCase.value, testCase.precision, got, testCase.want)
		}
	}
}

func TestDecimalArithmetic(t *testing.T) {
	t.Parallel()

	if got := Float.Add(0.1, 0.2); got != 0.3 {
		t.Fatalf("Float.Add(0.1, 0.2) = %v, want 0.3", got)
	}
	if got := Float.Mul(0.1, 0.2); got != 0.02 {
		t.Fatalf("Float.Mul(0.1, 0.2) = %v, want 0.02", got)
	}
	if got := Float.Divide(1, 3); got != 0.33 {
		t.Fatalf("Float.Divide(1, 3) = %v, want 0.33", got)
	}
	if got := Float.Divide(2, 3, 3); got != 0.667 {
		t.Fatalf("Float.Divide(2, 3, 3) = %v, want 0.667", got)
	}
	if got := Float.Divide(1, 0); got != 0 {
		t.Fatalf("Float.Divide(1, 0) = %v, want 0", got)
	}
	if got := Float.Round(1.235); got != 1.24 {
		t.Fatalf("Float.Round(1.235) = %v, want 1.24", got)
	}
	if got := Float.Round(125, -1); got != 130 {
		t.Fatalf("Float.Round(125, -1) = %v, want 130", got)
	}
	if got := Float.RoundInt64(125, -1); got != 130 {
		t.Fatalf("Float.RoundInt64(125, -1) = %v, want 130", got)
	}
}

func TestTruncateMatchesLegacyBehavior(t *testing.T) {
	t.Parallel()

	values := []float64{
		0,
		1.0 / 3.0,
		-1.0 / 3.0,
		1.9999999999999998,
		123.456789,
		-123.456789,
		0.000009,
		-0.000009,
		math.SmallestNonzeroFloat64,
		math.MaxFloat64,
	}
	for _, value := range values {
		for precision := 0; precision <= 8; precision++ {
			want := legacyTruncate(value, precision)
			if got := Float.Truncate(value, precision); got != want {
				t.Fatalf("Float.Truncate(%v, %d) = %v, legacy = %v", value, precision, got, want)
			}
		}
	}
}

// legacyTruncate 保留迁移前 FloatToDecimal 的有效输入行为，用于回归比较。
func legacyTruncate(value float64, precision int) float64 {
	text := strconv.FormatFloat(value, 'f', -1, 64)
	if text == "" {
		return 0
	}
	if precision >= len(text) {
		return value
	}
	parts := strings.Split(text, ".")
	if len(parts) < 2 || precision >= len(parts[1]) {
		return value
	}
	result, _ := strconv.ParseFloat(parts[0]+"."+parts[1][:precision], 64)
	return result
}
