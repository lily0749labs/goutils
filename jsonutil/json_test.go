package jsonutil

import "testing"

func TestMarshalString(t *testing.T) {
	t.Parallel()

	value := struct {
		Text string `json:"text"`
	}{Text: "<中文>"}

	got, err := JSON.MarshalString(value)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"text":"\u003c中文\u003e"}`; got != want {
		t.Fatalf("JSON.MarshalString() = %q, want %q", got, want)
	}

	got, err = JSON.MarshalStringNoEscapeHTML(value)
	if err != nil {
		t.Fatal(err)
	}
	if want := "{\"text\":\"<中文>\"}\n"; got != want {
		t.Fatalf("JSON.MarshalStringNoEscapeHTML() = %q, want %q", got, want)
	}
}

func TestMarshalStringError(t *testing.T) {
	t.Parallel()

	if got, err := JSON.MarshalString(func() {}); err == nil || got != "" {
		t.Fatalf("JSON.MarshalString(func) = %q, %v", got, err)
	}
	if got, err := JSON.MarshalStringNoEscapeHTML(func() {}); err == nil || got != "" {
		t.Fatalf("JSON.MarshalStringNoEscapeHTML(func) = %q, %v", got, err)
	}
}
