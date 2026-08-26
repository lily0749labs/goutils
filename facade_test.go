package goutils_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lily0749labs/goutils"
	anyutil "github.com/lily0749labs/goutils/anyto"
	cryptoutil "github.com/lily0749labs/goutils/crypto"
	idutil "github.com/lily0749labs/goutils/id"
	randutil "github.com/lily0749labs/goutils/rand"
	strutil "github.com/lily0749labs/goutils/strto"
	timeutil "github.com/lily0749labs/goutils/time"
	validutil "github.com/lily0749labs/goutils/valid"
)

func TestFacades(t *testing.T) {
	value, err := goutils.AnyTo.Int("42")
	if err != nil || value != 42 {
		t.Fatalf("AnyTo.Int() = %d, %v; want 42, nil", value, err)
	}

	if got := goutils.StrTo.Int("42"); got != 42 {
		t.Fatalf("StrTo.Int() = %d, want 42", got)
	}
	if got := goutils.StrTo.Ints([]string{"1", "2"}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("StrTo.Ints() = %v, want [1 2]", got)
	}

	if !goutils.Valid.Email("test@example.com") {
		t.Fatal("Valid.Email() = false, want true")
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

	if got := goutils.ID.Token(); !strings.Contains(got, "-") {
		t.Fatalf("ID.Token() = %q, want a UUID", got)
	}
	if got := goutils.ID.SnowflakeID(); got <= 0 {
		t.Fatalf("ID.SnowflakeID() = %d, want a positive ID", got)
	}

	if got, want := goutils.Time.CurrentLayout(), timeutil.Time.CurrentLayout(); got != want {
		t.Fatalf("Time.CurrentLayout() = %q, want %q", got, want)
	}
}

func TestLegacyFunctionsRemainCompatible(t *testing.T) {
	value, err := anyutil.AnyToInt("42")
	if err != nil || value != 42 {
		t.Fatalf("anyto.AnyToInt() = %d, %v; want 42, nil", value, err)
	}
	if got := strutil.StrToInt("42"); got != 42 {
		t.Fatalf("strto.StrToInt() = %d, want 42", got)
	}
	if !validutil.IsEmail("test@example.com") {
		t.Fatal("valid.IsEmail() = false, want true")
	}
	if got := cryptoutil.MD5("goutils"); got != goutils.Crypto.MD5("goutils") {
		t.Fatalf("crypto.MD5() = %q, want %q", got, goutils.Crypto.MD5("goutils"))
	}
	if got := randutil.Rand4(); len(got) != 4 {
		t.Fatalf("len(rand.Rand4()) = %d, want 4", len(got))
	}
	if got := idutil.GetToken(); !strings.Contains(got, "-") {
		t.Fatalf("id.GetToken() = %q, want a UUID", got)
	}
}
