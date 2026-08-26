package strto

import (
	"reflect"
	"testing"
)

func TestCommaInts(t *testing.T) {
	t.Parallel()

	if got := StrTo.CommaInts(""); got != nil {
		t.Fatalf("StrTo.CommaInts(empty) = %#v", got)
	}
	if got, want := StrTo.CommaInts("1,,bad,2,-3"), []int{1, 2, -3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("StrTo.CommaInts() = %#v, want %#v", got, want)
	}
}

func TestParseCommaInts(t *testing.T) {
	t.Parallel()

	if got := mustParseCommaInts(t, ""); got != nil {
		t.Fatalf("StrTo.ParseCommaInts(empty) = %#v", got)
	}
	if got, want := mustParseCommaInts(t, ",1, 2,-3,"), []int{1, 2, -3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("StrTo.ParseCommaInts() = %#v, want %#v", got, want)
	}
	if got, err := StrTo.ParseCommaInts("1,bad,2"); err == nil || got != nil {
		t.Fatalf("StrTo.ParseCommaInts(invalid) = %#v, %v", got, err)
	}
}

func mustParseCommaInts(t *testing.T, value string) []int {
	t.Helper()
	result, err := StrTo.ParseCommaInts(value)
	if err != nil {
		t.Fatal(err)
	}
	return result
}
