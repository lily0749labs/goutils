package stringutil

import "testing"

func TestEscapeSpecial(t *testing.T) {
	t.Parallel()

	if got, want := String.EscapeSpecial("abc-中文+$"), `abc\-\中\文\+\$`; got != want {
		t.Fatalf("String.EscapeSpecial() = %q, want %q", got, want)
	}
	if got := String.EscapeSpecial("abc123"); got != "abc123" {
		t.Fatalf("String.EscapeSpecial(plain) = %q", got)
	}
}
