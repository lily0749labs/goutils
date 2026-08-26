package goutils_test

import (
	"strings"
	"testing"

	"github.com/lily0749labs/goutils"
	timeutil "github.com/lily0749labs/goutils/time"
)

func TestFacades(t *testing.T) {
	value, err := goutils.AnyTo.AnyToInt("42")
	if err != nil || value != 42 {
		t.Fatalf("AnyTo.AnyToInt() = %d, %v; want 42, nil", value, err)
	}

	if got := goutils.StrTo.StrToInt("42"); got != 42 {
		t.Fatalf("StrTo.StrToInt() = %d, want 42", got)
	}

	if !goutils.Valid.IsEmail("test@example.com") {
		t.Fatal("Valid.IsEmail() = false, want true")
	}

	if got := goutils.Crypto.MD5("goutils"); len(got) != 32 {
		t.Fatalf("len(Crypto.MD5()) = %d, want 32", len(got))
	}

	const (
		key     = "1234567890123456"
		message = "goutils"
	)
	ciphertext, err := goutils.Crypto.Encrypt(key, message)
	if err != nil {
		t.Fatalf("Crypto.Encrypt() error = %v", err)
	}
	plaintext, err := goutils.Crypto.Decrypt(key, ciphertext)
	if err != nil || string(plaintext) != message {
		t.Fatalf("Crypto.Decrypt() = %q, %v; want %q, nil", plaintext, err, message)
	}

	if got := goutils.Rand.Rand4(); len(got) != 4 {
		t.Fatalf("len(Rand.Rand4()) = %d, want 4", len(got))
	}

	if got := goutils.ID.GetToken(); !strings.Contains(got, "-") {
		t.Fatalf("ID.GetToken() = %q, want a UUID", got)
	}

	if got, want := goutils.Time.CurrentLayout(), timeutil.Time.CurrentLayout(); got != want {
		t.Fatalf("Time.CurrentLayout() = %q, want %q", got, want)
	}
}
