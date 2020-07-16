// example03_test
package example03

import (
	"testing"
)

func TestUSTranslation(t *testing.T) {
	got := Greeting("George", "en-US")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestTranslation(t *testing.T) {
	got := translate("fr")
	want := "Bonjour"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
